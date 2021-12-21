-Projede cmd dizinine gittikten sonra "go run main.go" ile çalıştırdıktan sonra aşağıdaki endpointlere istek atılabilir;

  /getMemoryData/{key} =>(GET) redis de ki key datasının karşılığını döndürür.
  
  /setMemoryData => (POST) redis'e request de ki key,value objesine göre data yazar. Örnek parametre;
    {
        "key": "key2",
        "value": "value2"
    }
    
  /flushMemoryData => (DELETE) redis de ki tüm datayı siler.
  
-Uygulama ilk çalıştığında "tmp/data.json" dosyasında ki veriler redis'e yazılır.

-Dakika da bir çalışan cron ile redis de ki data "tmp/data.json" dosyasına yazılır.

-Redis'in lokalde 6379 portunda çalışır şekilde olması gerekmektedir. Config bu şekildedir. (127.0.0.1:6379)

-Endpoint response'ları ortak bir yapıdadır;

{
    "success": true,
    "errorMessage": "",
    "data": ""
}
