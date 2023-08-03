# Documentation for Snook IoT 

## Purpose
The purpose of this domain area is to connect the physical world to the digital world and making the other domain areas in the Snook project able to communicate in a simple and 
intuitive way with these devices and making use of their funcitonalites.

## Introduction 
This module of the Snook project is mainly built for enrolling IoT devices into the communication network. The purpose is to make it possible for
IoT devices to be apart of the WebSocket communication, to create value to and from the physical world. This is achieved mainly by showing or delivering data. 
At this time, the development of this domain area has been developed with a ESP32 Dev Kit C V4 microcontroller. However, the API:s and libraries used can be used for other Arduino / ESP32  boards
that support Wifi communication and Arduino software.

## Installation
The project is built with platform.io, so an installation of the platform.io extenstion in VS Code is the best and easiest way to get started developing. However, making use of the Arduino IDE is also an option,
but this requires a different project setup that is not implemented here.

After cloning the repo, secrets regarding Wifi connection needs to be placed inside a header file called "secrets.h" in the src folder. This needs to contain the following values:

#define SECRET_WIFI_SSID "your_wifi_ssid";  
#define SECRET_WIFI_PASSWORD "your_wifi_password";  
#define SERVER_IP "ip_of_snook_server";

## TODOS
Different modules for retrieving and sending data from IoT clients, showing specific data on specific devices, more functionalities, routines on specific devices, more fun stuff. 

## Parts used during development:

ESP32 Dev Kit C V4
https://www.amazon.se/AZDelivery-ESP32-NodeMCU-Development-Efterf%C3%B6ljarmodul/dp/B07Z83H831/ref=sr_1_2?keywords=esp32%2Bdev%2Bkit%2Bc%2Bv4&qid=1691059594&sprefix=dev%2Bkit%2Caps%2C92&sr=8-2&th=1

I2C LCD display with 
https://www.amazon.se/Freenove-Display-Compatible-Arduino-Raspberry/dp/B0B76Z83Y4/ref=sr_1_20?crid=1BKZEN243VODX&keywords=i2c&qid=1691059611&sprefix=i2c%2Caps%2C85&sr=8-20&th=1

## If you just like me is new to these things, here is some good reading:

### Development with microcontrollers
https://docs.platformio.org/en/latest/what-is-platformio.html
https://docs.espressif.com/projects/esp-idf/en/latest/esp32/hw-reference/esp32/get-started-devkitc.html

### I2C communication protocol (making it easier to develop lcd displays)
https://focuslcds.com/i2c-display-communication/
