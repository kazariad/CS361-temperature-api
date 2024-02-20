## Temperature API

The temperature API service consists of 2 **GET** endpoints which can be accessed with any standard HTTP client/library.
They do not require any parameters, special headers, or authentication. Both endpoints can be tested by opening them
directly in a web-browser, or with a utility such as **curl**.

> **http://74.208.168.118/temp/insight**
>- If successful, this endpoint will return status code **200** along with an insightful message containing some useful
   tip, news, or reminder.
>- The message will be encoded as UTF-8 plain text inside the response body.
>> curl --location --request GET 'http://74.208.168.118/temp/insight'
>
> ![](./insight.svg)

> **http://74.208.168.118/temp/version**
>- If successful, this endpoint will return status code **200** along with the temperature program's current version
   number.
>- The version number will be encoded as UTF-8 plain text inside the response body in standard semantic format (e.g.
   v1.3.4).
>> curl --location --request GET 'http://74.208.168.118/temp/version'
>
> ![](./version.svg)
