# Circle Matcher

Circle Matcher is a HTTP service that allows you validade and/or identify, through defined logic rules during segmentations creation, if your circles are able to find the users you want them to. 

There are two ways to validade segmentation through Circle Matcher:

1. **Default:** it is the manual validation where you will add up all the keys that defines the predetermined characteristic to a circle. _-- é a validação manual, em que você vai adicionando todas as chaves que definem as características pré-determinadas para_ círculo.   
2. **JSON:** it is a validation which goes straight to the JSON on your production environment and inserts a field on the paylod so it will make the try. _a validação na qual você vai direto no JSON do seu ambiente produtivo e insere no campo de payload para, em seguida, fazer o try._

If you pass an attribute that is out of the configured conditions on the circles, the system will return that specific user is in [**open sea**](https://docs.charlescd.io/principais-conceitos#mar-aberto), which means, it is in the general circle segmentation. 

## Circles identification

Se após a criação do círculo for necessária a utilização do Circle Matcher para testar suas regras de segmentação, você pode integrar nas suas aplicações o recurso Identify do módulo `charles-circle-matcher` para detectar os círculos aos quais o seu usuário pertence.

Por exemplo, dada a utilização dos seguintes parâmetros ao segmentar:

![](https://lh6.googleusercontent.com/q573-961WtpntVK8NfXXvPgzSPrxLwxjx3QXRqM3vBlHFM8nAoDkpn1KD26Zfw3_wJtjnhVldYcwRUUzhbveEvqJz6n16NQFkxi0S3hh8rk6Y7OUmWtnBOl_qJekzoymQ64mFF8k)

Ao realizar a requisição de identificação com as seguintes informações, círculos compatíveis serão retornados:

{% api-method method="get" host="https:" path="//api.charles-circle-matcher.com/identify" %}
{% api-method-summary %}
Identify
{% endapi-method-summary %}

{% api-method-description %}

{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="x-application-id" type="string" required=true %}
Workspace's ID
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}

{% api-method-body-parameters %}
{% api-method-parameter name="State" type="string" required=false %}
NY
{% endapi-method-parameter %}

{% api-method-parameter name="Job Title" type="string" required=false %}
Lawyer 
{% endapi-method-parameter %}

{% api-method-parameter name="Age" type="number" required=false %}
46 
{% endapi-method-parameter %}

{% api-method-parameter name="City" type="string" required=false %}
Stony Brook
{% endapi-method-parameter %}
{% endapi-method-body-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```
{ 
    "circles": [ 
    { 
    "id": "6577ae92-648c-11ea-bc55-0242ac130003", "name": "NY Lawyers" 
},  { 
    "id": "6577b112-648c-11ea-bc55-0242ac130003", 
    "name": "Stony Brook's Citizens" 
    } 
  ] 
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

Como no nosso exemplo existem círculos correspondentes com as informações sobre o usuário, o charles-circle-matcher está retornando uma lista com eles. Neste caso, dois círculos se encaixaram: NY Lawyers e Stony Brook’s Citizens.

Nessa requisição, apenas o parâmetro x-application-id é obrigatório. O body é totalmente flexível, porém vale lembrar que as chaves devem ter a mesma nomenclatura definida nas regras de segmentação do círculo. Veja no caso a seguir:

![](https://lh3.googleusercontent.com/FdPVIHDFeYJCkC_6Y1P3ZOBSqmNlGkl9q2_XyIayNKQo2Mp9IXBY7PzvpzW0Mej1P9Ox8AG12QiA1H0w5uozWP1UYWafcfwXLKBOf3G-ObIVoPHtYGOlWd5Ju01uLuScqtCn8qQ1)

O círculo “Stony Brook’s Citizens” foi criado para a identificar usuários que tenham como característica a chave “city” e o exato valor “Stony Brook”. Sendo assim, ele não estará na listagem ao realizar uma requisição para o Identify caso seja informado o body como no exemplo abaixo.

Outro ponto é que, quando o usuário não se enquadra em nenhuma segmentação, o sistema retorna indicando que ele se encaixa no “Mar Aberto”, isto é, em uma espécie de segmentação geral que inclui todos os usuários que estão fora de um círculo específico:

{% api-method method="get" host="https://" path="api.charles-circle-matcher.com/identify" %}
{% api-method-summary %}
Open Sea
{% endapi-method-summary %}

{% api-method-description %}

{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="x-application-id" type="string" required=true %}
Workspace's ID 
{% endapi-method-parameter %}
{% endapi-method-headers %}

{% api-method-body-parameters %}
{% api-method-parameter name="Age" type="number" required=false %}
46
{% endapi-method-parameter %}

{% api-method-parameter name="City" type="string" required=false %}
Stony Brook
{% endapi-method-parameter %}
{% endapi-method-body-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
List with all circles a user belongs to 
{% endapi-method-response-example-description %}

```
{
   "circles":  [
    {
       "id": 6577ae92-648c-11ea-bc55-0242ac130003", 
       "name": "NY Lawyers"
    },
    { 
       "id": 6577b1112-648c-11ea-bc55-0242ac30003",
       "name": "Stony Brook's Citizens"
    }
  ]
}   
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

Uma boa prática é realizar essa identificação sempre que o usuário faz login na aplicação. Entretanto, isso pode ser alterado de acordo com a necessidade da sua regra de negócio.
