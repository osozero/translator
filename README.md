###Command-line translation tool 

With the Translator tool, which uses the Yandex Translator service, you can quickly and freely translate from one language to another in the command line. You can save the API Key information that you can get for free from Yandex in the configuration file and start using Translator.

Usage: translator -t Hello -l en-tr

-t Text to be translated. **If text contains space, quotation marks must be used**

-l Translation direction, **default value is en-tr** which translates text from English to Turkish


####Use case 1:

>Translator -t "Hello world"

>Response:
>Hello world: [Merhaba dünya]

####Use case 2:

>Translator -t hello -l tr-es

>Response:
>Hello: [Hola]

Have fun :)
