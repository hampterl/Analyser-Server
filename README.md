# Analyser-Server


## HOW TO:
- Add Webhook to code -> preferably a discord webhook because the json format already inside is in discords format
- Create a .exe with no console (invisible exe)
- Open the app

## Description
The mouse will load for a bit ~5 sec and thats it, nothing visible will happen after that.  

Under the hood and not visible to the user, analyser will create a directory in Programm Data named Analyse Server, then inside of that 4 more named differently
one of them is called bin the exe will copy itself into there. After that it will copy a .lnk into the autostart and will continue to open itself
after the computer starts and the user logs in.

Deleting only works after closing the process in Task Manager and then you can delete the whole directory Analyse Server in Programm Data,
deleting the .exe in downloads or whatever other directory where you stored it bevore opening, will not work since the real copy is in Programm Data.

## Disclaimer
I the Author am not liable to any Damage/Harm/Data-loss caused by miss-use.  
This app is only for educational purpose and to prove that I understand the code/problem solving process of offensive security tool development.

## Author
**hampterl**
GitHub: [https://github.com/hampterl](https://github.com/hampterl)
