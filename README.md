# Winget
Winget is try to download, install and manage the application in your computer or in the web
<br/>
基本上Winget最重要的就幾個點 **search**, **list**, **download**, **install**, **upgrade**
<br/>

#### search
他是利用網路去各個地方去找相關的應用程式

#### list
他是利用自己電腦既有的應用程式取搜尋

#### download
從網路或是任何地方下載到**Downloads** folder裡面,但還沒安裝
<br/>
if you want to download the application into a specific folder, you can use **-d, --download-directory=/Path/you/want/to/download/application/**

#### install
安裝成應用程式

#### -q;--query
在各種的執行都可以利用他,只要用 **-q;--query=**就可能找到你想要的東西

<br/>
如果來源(source)是 **msstore**, 他就是從 **Microsoft Store**出來的,可能要錢!!


# Code of Testing Winget with golang

### wingetList -- wingetListMonikerSource() function
Want to search winget package in the web and find that they might have the " " so you might find it and download the specific pakcage
<br/>
Now the suitable is function **wingetListMonikerSource**, it can search the application in the web


### wingetDownload -- wingetDownload.go function

