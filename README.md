## Usage
Global variable sets up the html tags whitelist which is a single selector:
```
div,span,hr,p,a,img,strong,i,b,em,li,ul,h6,h5,h4,h3,h2,h1,blockquote,html
```
It is applied using `.Not` to fetch all the children of the given html fragment using `goquery`.

### Title regexes
```
([\{\(\[][^()\[\]]*[\)\]\}])*
\*{2,}
(.)1{3,}
```
### Api
`POST` where the body is the raw content to parse. Default port `8002`
```
/v1/body
/v1/title
```

