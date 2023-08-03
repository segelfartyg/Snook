#include <Arduino.h>
#include <WiFi.h>
#include <Adafruit_SSD1306.h>
#include <Adafruit_Sensor.h>
#include <LiquidCrystal_I2C.h>
#include <WebSocketsClient.h>
#include <secrets.h>

LiquidCrystal_I2C lcd(0x27, 16, 2);

const char *ssid_router = SECRET_WIFI_SSID;
const char *password_router = SECRET_WIFI_PASSWORD;
char server[] = SERVER_IP;

WiFiClient httpclient;
WebSocketsClient webSocket;

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

void setupWifi()
{

  WiFi.disconnect();

  Serial.println("BEGINNING WIFI CONNECTION");
  WiFi.begin(ssid_router, password_router);

  while (WiFi.status() != WL_CONNECTED)
  {
    Serial.println(WiFi.status());
    Serial.println("Connecting...");
    delay(500);
  }
  Serial.println("CONNECTED IP:");
  Serial.println(WiFi.localIP());
}

String dataReader(WiFiClient _httpclient)
{

  String response;
  bool responseReceived = false;
  bool timeOut = false;
  int retries = 0;

  while (!responseReceived && !timeOut)
  {

    while (_httpclient.available())
    {
      char c = _httpclient.read();
      response += c;
    }

    if (response.length() > 0)
    {
      responseReceived = true;
    }
    else
    {
      if (retries < 5)
      {
        delay(1000);
        retries++;
      }
      else
      {
        response = "INTERNAL READING ERROR (001)";
        ESP.restart();
        return response;
      }
    }
  }

  Serial.println(response);
  return response;
}

void hexdump(const void *mem, uint32_t len, uint8_t cols = 16) {
	const uint8_t* src = (const uint8_t*) mem;
	Serial.printf("\n[HEXDUMP] Address: 0x%08X len: 0x%X (%d)", (ptrdiff_t)src, len, len);
	for(uint32_t i = 0; i < len; i++) {
		if(i % cols == 0) {
			Serial.printf("\n[0x%08X] 0x%08X: ", (ptrdiff_t)src, i);
		}
		Serial.printf("%02X ", *src);
		src++;
	}
	Serial.printf("\n");
}

void showLCDMessage(uint8_t * message){

  lcd.clear();
  lcd.setCursor(0, 0);
  lcd.print(String((char *)message));

}

void webSocketEvent(WStype_t type, uint8_t * payload, size_t length){

	switch(type) {
		case WStype_DISCONNECTED:
			Serial.println("[WSc] Disconnected!\n");
			break;
		case WStype_CONNECTED:
			Serial.printf("[WSc] Connected to url: %s\n", payload);

			// send message to server when Connected
			webSocket.sendTXT("Connected");
			break;
		case WStype_TEXT:
			Serial.printf("[WSc] get text: %s\n", payload);
      showLCDMessage(payload);

			// send message to server
			// webSocket.sendTXT("message here");
			break;
		case WStype_BIN:
			Serial.printf("[WSc] get binary length: %u\n", length);
			hexdump(payload, length);

			// send data to server
			// webSocket.sendBIN(payload, length);
			break;
		case WStype_ERROR:			
		case WStype_FRAGMENT_TEXT_START:
		case WStype_FRAGMENT_BIN_START:
		case WStype_FRAGMENT:
		case WStype_FRAGMENT_FIN:
			break;
	}

}


void setup()
{

  webSocket.begin("192.168.1.138", 80, "/ws");

  webSocket.onEvent(webSocketEvent);

  // try ever 5000 again if connection has failed
	webSocket.setReconnectInterval(5000);


  Serial.begin(9600);

  if (!i2CAddrTest(0x27))
  {
    lcd = LiquidCrystal_I2C(0x3F, 16, 2);
  }

  setupWifi();

  lcd.init();
  lcd.backlight();
  lcd.setCursor(0, 0);
}

void loop()
{
  webSocket.loop();
}
