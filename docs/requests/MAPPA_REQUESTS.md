# Mappa Requests

## Escotista

### Escotista Exists

```
GET /api/escotistas/44471/exists HTTP/1.1
accept: application/json
cache-control: no-cache
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
Content-Length: 15
ETag: W/"f-lWSD52foMX4qLT82tweJE091S9Q"
Date: Tue, 03 Sep 2019 12:10:03 GMT

{"exists":true}
```

### Escotistas Sessões

```
GET /api/escotistas/44471/secoes HTTP/1.1
accept: application/json
cache-control: no-cache
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
Content-Length: 96
ETag: W/"60-xnUN2znJVSwgSKv6n1/5iVHcdyQ"
Date: Tue, 03 Sep 2019 12:10:03 GMT

[{"codigo":1424,"nome":"ALCATÉIA 1 ","codigoTipoSecao":1,"codigoGrupo":32,"codigoRegiao":"SC"}]
```

### Escotistas Sessões Equipes

```
GET /api/escotistas/44471/secoes/1424/equipes?filter={%22include%22:%20%22associados%22} HTTP/1.1
accept: application/json
cache-control: no-cache
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
ETag: W/"25d5-Me6LkcM5wrZYlW4pCEVM8mZIY1g"
Content-Encoding: gzip
Date: Tue, 03 Sep 2019 12:10:03 GMT
Transfer-Encoding: chunked

[
	{
		"codigo": 5649,
		"nome": "Matilha Amarela (Pedro)",
		"codigoSecao": 1424,
		"codigoLider": 992696,
		"codigoViceLider": 1102053,
		"associados": [
			{
				"codigo": 992696,
				"nome": "BENICIO AMADEU WOLFF PEREIRA",
				"codigoFoto": 33689,
				"codigoEquipe": 5649,
				"username": 1263323,
				"numeroDigito": 2,
				"dataNascimento": "Tue Sep 13 2011 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2022-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1102051,
				"nome": "BERNARDO GROTMANN MORITZ",
				"codigoFoto": 32276,
				"codigoEquipe": 5649,
				"username": 1439300,
				"numeroDigito": 0,
				"dataNascimento": "Wed Feb 05 2014 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2022-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1102053,
				"nome": "GABRIEL FRANZ",
				"codigoFoto": null,
				"codigoEquipe": 5649,
				"username": 1439302,
				"numeroDigito": 6,
				"dataNascimento": "Thu Aug 22 2013 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2021-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1113108,
				"nome": "LOUISE VICTORIA BETHE FURLAN",
				"codigoFoto": null,
				"codigoEquipe": 5649,
				"username": 1449395,
				"numeroDigito": 0,
				"dataNascimento": "Fri Feb 07 2014 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2021-01-01T00:00:00.000Z",
				"nomeAbreviado": "LOUISE V. BETHE FURLAN",
				"sexo": "F",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1120013,
				"nome": "SAMUEL BRANDÃO GARROTI",
				"codigoFoto": null,
				"codigoEquipe": 5649,
				"username": 1456367,
				"numeroDigito": 3,
				"dataNascimento": "Fri Aug 16 2013 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2021-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1137416,
				"nome": "ANDREAS MIGUEL HEINZEN",
				"codigoFoto": null,
				"codigoEquipe": 5649,
				"username": 1472150,
				"numeroDigito": 3,
				"dataNascimento": "Wed Jan 22 2014 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2022-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			}
		]
	},
	{
		"codigo": 5650,
		"nome": "Matilha Cinza (Guionardo)",
		"codigoSecao": 1424,
		"codigoLider": 1047778,
		"codigoViceLider": 1104368,
		"associados": [
			{
				"codigo": 1047778,
				"nome": "MIGUEL BARRETO HADLICH",
				"codigoFoto": 22805,
				"codigoEquipe": 5650,
				"username": 1371138,
				"numeroDigito": 5,
				"dataNascimento": "Sat May 26 2012 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2022-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1102061,
				"nome": "ISAAC GUSTAVO COSTA",
				"codigoFoto": 32437,
				"codigoEquipe": 5650,
				"username": 1439306,
				"numeroDigito": 9,
				"dataNascimento": "Thu Aug 21 2014 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2022-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1104368,
				"nome": "LETICIA FORTUNATO",
				"codigoFoto": null,
				"codigoEquipe": 5650,
				"username": 1441479,
				"numeroDigito": 1,
				"dataNascimento": "Mon Mar 19 2012 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2021-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "F",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1113131,
				"nome": "MATHEUS HENRIQUE ERDMANN DE NASCIMENTO",
				"codigoFoto": null,
				"codigoEquipe": 5650,
				"username": 1454116,
				"numeroDigito": 5,
				"dataNascimento": "Sat Mar 28 2015 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2022-01-01T00:00:00.000Z",
				"nomeAbreviado": "MATHEUS H. E. DE NASCIMENTO",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1114599,
				"nome": "LUIZ OTAVIO SAUER WALKER",
				"codigoFoto": null,
				"codigoEquipe": 5650,
				"username": 1450940,
				"numeroDigito": 7,
				"dataNascimento": "Wed Jan 16 2013 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2021-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1128140,
				"nome": "MANUELA JACOMEL",
				"codigoFoto": null,
				"codigoEquipe": 5650,
				"username": 1463956,
				"numeroDigito": 4,
				"dataNascimento": "Thu Apr 18 2013 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2022-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "F",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			}
		]
	},
	{
		"codigo": 5651,
		"nome": "Matilha Marrom (Josi)",
		"codigoSecao": 1424,
		"codigoLider": 1101641,
		"codigoViceLider": 1088573,
		"associados": [
			{
				"codigo": 1063368,
				"nome": "CAIO BATISTA DE OLIVEIRA",
				"codigoFoto": null,
				"codigoEquipe": 5651,
				"username": 1398759,
				"numeroDigito": 3,
				"dataNascimento": "Tue Sep 04 2012 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2022-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1088573,
				"nome": "JOAQUIM ZIMMERMANN DE ASSIS",
				"codigoFoto": null,
				"codigoEquipe": 5651,
				"username": 1425184,
				"numeroDigito": 1,
				"dataNascimento": "Mon Jul 15 2013 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2021-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1101641,
				"nome": "AMANDA VON HOHENDORFF MAAS DE FAVERI",
				"codigoFoto": 32346,
				"codigoEquipe": 5651,
				"username": 1438959,
				"numeroDigito": 2,
				"dataNascimento": "Thu Jan 02 2014 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2022-01-01T00:00:00.000Z",
				"nomeAbreviado": "AMANDA V H M DE FAVERI",
				"sexo": "F",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1103478,
				"nome": "EMANUELA DA SILVA BULSONI",
				"codigoFoto": null,
				"codigoEquipe": 5651,
				"username": 1440687,
				"numeroDigito": 0,
				"dataNascimento": "Sun Jul 21 2013 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2021-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "F",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1128118,
				"nome": "PEDRO HENRIQUE CARVALHO ARAUJO DA SILVA",
				"codigoFoto": null,
				"codigoEquipe": 5651,
				"username": 1463954,
				"numeroDigito": 8,
				"dataNascimento": "Mon Jan 27 2014 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2022-01-01T00:00:00.000Z",
				"nomeAbreviado": "PEDRO HENRIQUE C A DA SILVA",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			}
		]
	},
	{
		"codigo": 5652,
		"nome": "Matilha Vermelha (Rita)",
		"codigoSecao": 1424,
		"codigoLider": 1025980,
		"codigoViceLider": 1105848,
		"associados": [
			{
				"codigo": 1025980,
				"nome": "ALICE SCHMITT BERNARDES",
				"codigoFoto": null,
				"codigoEquipe": 5652,
				"username": 1318116,
				"numeroDigito": 5,
				"dataNascimento": "Fri May 25 2012 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2022-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "F",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1103477,
				"nome": "BENJAMIN DOS SANTOS FURLAN",
				"codigoFoto": null,
				"codigoEquipe": 5652,
				"username": 1440686,
				"numeroDigito": 1,
				"dataNascimento": "Fri Feb 07 2014 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2021-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1105848,
				"nome": "BETINA SCHATZ SANDRINI DE CASTRO",
				"codigoFoto": null,
				"codigoEquipe": 5652,
				"username": 1442767,
				"numeroDigito": 2,
				"dataNascimento": "Fri Jan 04 2013 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2022-01-01T00:00:00.000Z",
				"nomeAbreviado": "BETINA S. S. DE CASTRO",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1107587,
				"nome": "LARA DOS SANTOS KUFKY",
				"codigoFoto": null,
				"codigoEquipe": 5652,
				"username": 1444661,
				"numeroDigito": 8,
				"dataNascimento": "Wed Jan 14 2015 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2021-01-01T00:00:00.000Z",
				"nomeAbreviado": "",
				"sexo": "F",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1113127,
				"nome": "THOMAS AUGUSTO TECHENTIN PACHECO",
				"codigoFoto": null,
				"codigoEquipe": 5652,
				"username": 1450935,
				"numeroDigito": 0,
				"dataNascimento": "Tue Apr 07 2015 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2021-01-01T00:00:00.000Z",
				"nomeAbreviado": "THOMAS A. T. PACHECO",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			},
			{
				"codigo": 1137421,
				"nome": "IAN BERNARDO MAÇANEIRO DA SILVA",
				"codigoFoto": null,
				"codigoEquipe": 5652,
				"username": 1472152,
				"numeroDigito": 0,
				"dataNascimento": "Tue Jul 02 2013 00:00:00 GMT+0000 (Coordinated Universal Time)",
				"dataValidade": "2022-01-01T00:00:00.000Z",
				"nomeAbreviado": "IAN B MAÇANEIRO DA SILVA",
				"sexo": "M",
				"codigoRamo": 2,
				"codigoCategoria": 1,
				"codigoSegundaCategoria": 0,
				"codigoTerceiraCategoria": 0,
				"linhaFormacao": null,
				"codigoRamoAdulto": null,
				"dataAcompanhamento": null
			}
		]
	}
]
```


### Escotista

```
GET /api/escotistas/44471 HTTP/1.1
accept: application/json
cache-control: no-cache
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
Content-Length: 176
ETag: W/"b0-3m7iETbzdMfmtMXn9nvLMqoDTsk"
Date: Tue, 03 Sep 2019 12:10:03 GMT

{"codigo":44471,"codigoAssociado":null,"username":"A 1 - Cinza","nomeCompleto":"Alcatéia 1 - Matilha Cinza","ativo":"S","codigoGrupo":32,"codigoRegiao":"SC","codigoFoto":null}
```


### Escotista Favoritos

```
GET /api/escotistas/44471/favoritos HTTP/1.1
accept: application/json
cache-control: no-cache
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
Content-Length: 2
ETag: W/"2-l9Fw4VUO7kr8CvBlt4zaMCqXZ0w"
Date: Tue, 03 Sep 2019 12:10:04 GMT

[]
```

### Escotista Codigo Foto (fields)

```
GET /api/escotistas/44471?filter[fields][codigoFoto]=true HTTP/1.1
accept: application/json
cache-control: no-cache
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
Content-Length: 19
ETag: W/"13-WVvTscFsq/L9Vu3c9f1i7S/f5Bs"
Date: Tue, 03 Sep 2019 12:10:04 GMT

{"codigoFoto":null}
```

## Escotista Informativos

```
GET /api/escotistas/44471/informativos HTTP/1.1
accept: application/json
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
Content-Length: 2
ETag: W/"2-l9Fw4VUO7kr8CvBlt4zaMCqXZ0w"
Date: Tue, 03 Sep 2019 12:10:04 GMT

[]
```

## Associados

### Associados Missões

GET /api/associados-missoes?filter={%22where%22:%20{%22codigoAssociado%22:%20null},%20%22include%22:%20%22missao%22,%20%22order%22:%20%22dataPlano%20DESC%22,%20%22limit%22:%2010,%20%22skip%22:%200%20} HTTP/1.1
accept: application/json
cache-control: no-cache
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
Content-Length: 2
ETag: W/"2-l9Fw4VUO7kr8CvBlt4zaMCqXZ0w"
Date: Tue, 03 Sep 2019 12:10:03 GMT

[]

### Associados Conquistas

```
GET /api/associado-conquistas/v2/updates?dataHoraUltimaAtualizacao=2019-08-31T13:48:22-00:00&codigoSecao=1424 HTTP/1.1
accept: application/json
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
Content-Length: 52
ETag: W/"34-F14SoK+i0DRvHVc71P3rSx7CVHU"
Date: Tue, 03 Sep 2019 12:10:10 GMT

{"dataHora":"2019-09-03T09:10:10-00:00","values":[]}
```

## Imagens

```
GET /api/imagens/8502 HTTP/1.1
accept: application/json
cache-control: no-cache
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
ETag: W/"180cb-bCCPUp/vJXD+hRCJe1AfKodJfBo"
Content-Encoding: gzip
Date: Tue, 03 Sep 2019 12:10:03 GMT
Transfer-Encoding: chunked

{ "imagem": "base64",
"id":8502}
```

## Avaliações

```
GET /api/avaliacoes/getAvaliacoes?codigoAssociado=null HTTP/1.1
accept: application/json
cache-control: no-cache
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
Content-Length: 2
ETag: W/"2-vyGp6PvFo4RvsFtPoIWeCReyIC8"
Date: Tue, 03 Sep 2019 12:10:04 GMT

{}
```

## Marcações updates

```
GET /api/marcacoes/v2/updates?dataHoraUltimaAtualizacao=2019-08-31T13:48:23-00:00&codigoSecao=1424 HTTP/1.1
accept: application/json
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
Content-Length: 52
ETag: W/"34-4oZ2S0qZSMKqQJmWoINOb/vpFIE"
Date: Tue, 03 Sep 2019 12:10:11 GMT

{"dataHora":"2019-09-03T09:10:11-00:00","values":[
    {"codigoAtividade":36,
    "codigoAssociado":102541,
    "dataAtividade":"2020-09-05T00:00:00.000Z",
    "dataStatusEscotista":"2020-09-05T17:43:53.000Z",
    "dataHoraAtualizacao":"2020-09-05T18:41:16.000Z",
    "codigoUltimoEscotista": 51121,
    "segmento":"PROGRESSAO_PATATENRA_SALTADOR"},
]}
```

## Marcações adultos updates

```
GET /api/marcacoesAdultos/updates?dataHoraUltimaAtualizacao=2019-08-31T13:48:23-00:00&codigoEscotista=44471 HTTP/1.1
accept: application/json
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
Content-Length: 52
ETag: W/"34-4oZ2S0qZSMKqQJmWoINOb/vpFIE"
Date: Tue, 03 Sep 2019 12:10:11 GMT

{"dataHora":"2019-09-03T09:10:11-00:00","values":[]}
```

## Grupos

```
GET /api/grupos?filter={%22where%22:%20{%22codigo%22:%2032,%20%22codigoRegiao%22:%20%22SC%22}} HTTP/1.1
accept: application/json
cache-control: no-cache
authorization: xiNSSgjDgSmWGLMbmVjOvQkrxXomAfQgiimNX40dCZBWnPiER2YnhjBphr2tHNC3
Host: mappa.escoteiros.org.br
Connection: Keep-Alive
Accept-Encoding: gzip
Cookie: __cfduid=d7a426e62ee0e95ca17d67fd7ae11617e1560788037
User-Agent: okhttp/3.4.1

HTTP/1.1 200 OK
Vary: Origin, Accept-Encoding
Access-Control-Allow-Credentials: true
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Download-Options: noopen
X-Content-Type-Options: nosniff
Content-Type: application/json; charset=utf-8
Content-Length: 84
ETag: W/"54-aWQn4djk3N5KQAOqWOqMF6qdqYc"
Date: Tue, 03 Sep 2019 12:10:03 GMT

[{"codigo":32,"codigoRegiao":"SC","nome":"LEÕES DE BLUMENAU","codigoModalidade":1}]
```

