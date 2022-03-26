## REST API Principles

###The Six Principles / Constraints
####1.Client-server
Client ve server birbirinden bağımsız olmalıdır. İstemci istediği bilgiyi sadece istenen kaynağın URI'sından ulaşır.
Başka hiç bir şekilde server ile etkilemişe geçemez.
Benzer şekilde server tarafı, clienttan gelen HTTP isteğine verileri iletmek dışında değiştirilmemelidir.


####2.Stateless
Clienttan server'a gelen her istek, isteği anlaman için gerekli tüm bilgileri içermelidir.
Server tarafında bir şey tutulmaz.

####3.Cacheable
Mümkün olduğunda, kaynaklar client veya server tarafında önbelleğe alınabilir olmalıdır.
Server yanıtlarının ayrıca, teslim edilen kaynak için önbelleğe almaya izin verilip verilmediği hakkında bilgi içermesi gerekir.
Amaç, sunucu tarafında ölçeklenebilirliği artırırken, istemci tarafında performansı iyileştirmektir.

####4.Layered System:
Bir client normalde doğrudan son sunucuya mı yoksa yol boyunca bir aracıya mı bağlı olduğunu söyleyemez.
Aracı serverlar, yük dengelemeyi etkinleştirerek ve paylaşılan önbellekler sağlayarak sistem ölçeklenebilirliğini iyileştirebilir. Katmanlar ayrıca güvenlik ilkelerini zorunlu kılabilir.

####5.Code-on-Demand(optional)
Bazı durumlarda response executable code barındırabilir. Bu gibi durumlarda code on-demand run edilmelidir.

####6.Uniform Interface
Servisin hangi endpointine istek atarsanız atın tüm cevaplar bir structure ile dönmelidir.
Tüm hata mesajlarının ve tüm başarılı isteklerin aynı yapıya sahip olmadı gerekir.

Source:
https://r.bluethl.net/how-to-design-better-apis