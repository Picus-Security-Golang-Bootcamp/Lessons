package pagination

import (
	"bytes"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		tag                                                                     string
		page, pageSize, total                                                   int
		expectedPage, expectedPageSize, expectedTotal, pageCount, offset, limit int
	}{
		// varying page
		{"t1", 1, 20, 50, 1, 20, 50, 3, 0, 20},
		{"t2", 2, 20, 50, 2, 20, 50, 3, 20, 20},
		{"t3", 3, 20, 50, 3, 20, 50, 3, 40, 20},
		{"t4", 4, 20, 50, 3, 20, 50, 3, 40, 20},
		{"t5", 0, 20, 50, 1, 20, 50, 3, 0, 20},

		// varying pageSize
		{"t6", 1, 0, 50, 1, 100, 50, 1, 0, 100},
		{"t7", 1, -1, 50, 1, 100, 50, 1, 0, 100},
		{"t8", 1, 100, 50, 1, 100, 50, 1, 0, 100},
		{"t9", 1, 1001, 50, 1, 1000, 50, 1, 0, 1000},

		// varying total
		{"t10", 1, 20, 0, 1, 20, 0, 0, 0, 20},
		{"t11", 1, 20, -1, 1, 20, -1, -1, 0, 20},
	}

	for _, test := range tests {
		p := New(test.page, test.pageSize, test.total)
		if test.expectedPage != p.Page {
			t.Errorf("Test failed [%s] : %v was given, %v want, %v got", test.tag, test.page, test.expectedPage, p.Page)
		}
		if test.expectedPageSize != p.PageSize {
			t.Errorf("Test failed [%s] : %v was given, %v want, %v got", test.tag, test.pageSize, test.expectedPageSize, p.PageSize)
		}

		if test.expectedPageSize != p.PageSize {
			t.Errorf("Test failed [%s] : %v was given, %v want, %v got", test.tag, test.pageSize, test.expectedPageSize, p.PageSize)
		}
		if test.expectedTotal != p.TotalCount {
			t.Errorf("Test failed [%s] : %v was given, %v want, %v got", test.tag, test.total, test.expectedTotal, p.TotalCount)
		}
		if test.pageCount != p.PageCount {
			t.Errorf("Test failed [%s] : %v want, %v got", test.tag, test.pageCount, p.PageCount)
		}
		if test.offset != p.Offset() {
			t.Errorf("Test failed [%s] : %v want, %v got", test.tag, test.offset, p.Offset())
		}
		if test.limit != p.Limit() {
			t.Errorf("Test failed [%s] : %v want, %v got", test.tag, test.limit, p.Limit())
		}
	}
}

func TestPages_BuildLinkHeader(t *testing.T) {
	baseURL := "/tokens"
	defaultPageSize := 10
	tests := []struct {
		tag                   string
		page, pageSize, total int
		header                string
	}{
		{"t1", 1, 20, 50, "</tokens?page=2&pageSize=20>; rel=\"next\", </tokens?page=3&pageSize=20>; rel=\"last\""},
		{"t2", 2, 20, 50, "</tokens?page=1&pageSize=20>; rel=\"first\", </tokens?page=1&pageSize=20>; rel=\"prev\", </tokens?page=3&pageSize=20>; rel=\"next\", </tokens?page=3&pageSize=20>; rel=\"last\""},
		{"t3", 3, 20, 50, "</tokens?page=1&pageSize=20>; rel=\"first\", </tokens?page=2&pageSize=20>; rel=\"prev\""},
		{"t4", 0, 20, 50, "</tokens?page=2&pageSize=20>; rel=\"next\", </tokens?page=3&pageSize=20>; rel=\"last\""},
		{"t5", 4, 20, 50, "</tokens?page=1&pageSize=20>; rel=\"first\", </tokens?page=2&pageSize=20>; rel=\"prev\""},
		{"t6", 1, 20, 0, ""},
		{"t7", 4, 20, -1, "</tokens?page=1&pageSize=20>; rel=\"first\", </tokens?page=3&pageSize=20>; rel=\"prev\", </tokens?page=5&pageSize=20>; rel=\"next\""},
	}
	for _, test := range tests {
		p := New(test.page, test.pageSize, test.total)
		got := p.BuildLinkHeader(baseURL, defaultPageSize)
		if test.header != got {
			t.Errorf("Test failed [%s] : %s want %s got", test.tag, test.header, got)
		}
	}

	baseURL = "/tokens?from=10"
	p := New(1, 20, 50)
	want := "</tokens?from=10&page=2&pageSize=20>; rel=\"next\", </tokens?from=10&page=3&pageSize=20>; rel=\"last\""
	got := p.BuildLinkHeader(baseURL, defaultPageSize)
	if want != got {
		t.Errorf("Test failed : %s want %s got", want, got)
	}

}

func Test_parseInt(t *testing.T) {
	type args struct {
		value        string
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"123", 100}, 123},
		{"t2", args{"", 100}, 100},
		{"t3", args{"a", 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInt(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("parseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com?page=2&pageSize=20", bytes.NewBufferString(""))
	p := NewFromRequest(req, 100)

	if 2 != p.Page {
		t.Errorf("Test failed NewFromRequest() [Page] = %v, want %v", p.Page, 2)
	}
	if 20 != p.PageSize {
		t.Errorf("Test failed NewFromRequest() [PageSize] = %v, want %v", p.PageSize, 20)
	}
	if 100 != p.TotalCount {
		t.Errorf("Test failed NewFromRequest() [TotalCount] = %v, want %v", p.TotalCount, 100)
	}

	if 5 != p.PageCount {
		t.Errorf("Test failed NewFromRequest() [PageCount] = %v, want %v", p.PageCount, 5)
	}
}
