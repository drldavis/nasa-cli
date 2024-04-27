# nasa-cli  
#### This application uses the Nasa API to allow users to obtain a weather report for mars, and to view data on solar flares for a given time range.

## Installation  
Clone this project  
`git clone https://github.com/drldavis/nasa-cli.git`  
Open the folder  
`cd nasa-cli`  
### Run the Mars Weather Report
##### Note: This application has been compiled to multiple executables for you to use based on your operating system. For the purpose of these instructions, I will use the `./nasa-mac-intel` executable.
`./nasa-mac-intel -mars`
#### Important
At the time of writing, the Mars Weather API does not currently have enough data to provide a weather report. The [Nasa API docs](https://api.nasa.gov/#:~:text=THIS%20SERVICE%20HAS%20SIGNIFICANT%20MISSING%20DATA%20DUE%20TO%20INSIGHT%20NEEDING%20TO%20MANAGE%20POWER%20USE%3APlease%20check%20out%20the%20seasonal%20weather%20report%20plot%20for%20an%20illustration%20of%20missing%20data%20and%20read%20this%20article%20about%20how%20dust%20and%20distance%20from%20the%20sun%20affect%20Insight%27s%20power%20situation.) warn that "THIS SERVICE HAS SIGNIFICANT MISSING DATA DUE TO INSIGHT NEEDING TO MANAGE POWER USE". If running `./nasa-mac-intel -mars` returns the following error: `error occured while getting mars weather: nasa does not currently have enough data to provide an accurate weather report on mars`, you will not be able to receive a weather report at this time. Try again later to see if Nasa has updated their data.  
### Run the Solar Flares Report
1. `./nasa-mac-intel -solar`  
2. Enter a start date when prompted
3. Enter an end date when prompted
4. View results
