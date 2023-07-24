#include <Arduino.h>
#include <WiFi.h>
#include <Adafruit_SSD1306.h>
#include <Adafruit_Sensor.h>
#include <LiquidCrystal_I2C.h>

LiquidCrystal_I2C lcd(0x27, 16, 2);

const char *ssid_router = "YourWifiName";
const char *password_router = "YourWifiPassword";
char server[] = "YourServerAddress";
WiFiClient httpclient;

bool i2CAddrTest(uint8_t addr)
{

  Wire.begin();
  Wire.beginTransmission(addr);
  if (Wire.endTransmission() == 0)
  {
    return true;
  }
  return false;
}


String dataReader(WiFiClient _httpclient){

  String response;
  bool responseReceived = false;

  while(!responseReceived){

    while(_httpclient.available()){
    char c = _httpclient.read();
    response += c;
    } 

    if(response.length() > 0){
      responseReceived = true;
    }
  }

  Serial.println(response);
  return response;

}

void setup()
{
  Serial.begin(9600);

  if (!i2CAddrTest(0x27))
  {
    lcd = LiquidCrystal_I2C(0x3F, 16, 2);
  }

  lcd.init();
  lcd.backlight();
  lcd.setCursor(0, 0);

  Serial.println("BEGINNING WIFI CONNECTION");
  WiFi.begin(ssid_router, password_router);

  while (WiFi.status() != WL_CONNECTED)
  {
    Serial.println("Connecting...");
    delay(500);
  }
  Serial.println("CONNECTED IP:");
  Serial.println(WiFi.localIP());
}

void loop()
{
  Serial.print("Making http req");
  if (httpclient.connect(server, 80))
  {
    httpclient.println("GET /ping HTTP/1.1");
    httpclient.println("Host: YourServerAddress");
    httpclient.println("Connection: close");
    httpclient.println();
  }


  String response = dataReader(httpclient);

    char *responseArray = new char[response.length()];
    strcpy(responseArray, response.c_str());
    String body = strstr(responseArray, "\n\r\n");
    String output;

    for (int i = 0; i < body.length(); i++)
    {

      if(body[i] != '\n' && body[i] != '\r')
      {
        output += body[i];
      }
    }

    lcd.clear();
    lcd.setCursor(0, 0);
    lcd.print(output);
  

  delay(2000);
}
