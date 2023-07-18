const express = require('express');
const axios = require('axios');
const Spotify= require('spotify-web-api-node');
const {SPOTIFY_CLIENT_ID, SPOTIFY_CLIENT_SECRET,API} = require('./config.js');
const weatherCodes=require('./models/weatherCodes.js');
const playlists=require('./models/playList.js');

require('dotenv').config()

const app = express()
const port = 5000

const api=API

/////////////////////////////spotify api/////////////////////////////////////

const spotifyApi = new Spotify({
  clientId: SPOTIFY_CLIENT_ID,
  clientSecret: SPOTIFY_CLIENT_SECRET
});


//using spotify web api node where clienrCredentialsGrant is used to get the access token from the cliend side 
async function getAccessToken() {
    try {
        const data = await spotifyApi.clientCredentialsGrant();
        const accessToken = data.body.access_token;
        console.log(accessToken);
        spotifyApi.setAccessToken(accessToken);
        return accessToken;        
  
    } catch (error) {
        console.error('Error authenticating with Spotify API:', error);
        throw new Error('Failed to authenticate with Spotify API');
    }
  }

  getAccessToken()
  .then(token => {
     console.log(token)
     getPlaylist('37i9dQZF1DX1gRalH1mWrP',token);
  })
  .catch(error => {
      console.error(error);
  });

/////////////////////////////tomoorrow api/////////////////////////////////////



app.use(express.json());

async function weather(location){
    try{
        const response = await axios.get(`https://api.tomorrow.io/v4/timelines?location=${location}&fields=temperature,precipitationIntensity,weatherCode&apikey=${api}`);
    //fields can be added or removed as per the requirement
    const weatherData = response.data.data.timelines[0].intervals[0].values;
    const temperature = weatherData.temperature;
    //returns the weather code from here 
    const weatherCode=  weatherData.weatherCode;
    //check the weather code and return the weather description
    const weatherDescription=weatherCodes[0][weatherCode];
    console.log(weatherDescription);
    console.log(temperature)
    console.log(location)
    console.log(weatherData);
    return weatherDescription;
    }
    catch(error){
        console.log(error);
        res.status(500).send('Some error occured');
    }
}
////////////////////////////////////weather code call/////////////////////////////////////

function playList(weatherDescription){
    if(weatherDescription.includes('Clear')||weatherDescription.includes('Sunny'))
    return playlists[0].sunny;
    else if(weatherDescription.includes('Rain')||weatherDescription.includes('Drizzle'))
    return playlists[0].rainy;
    else if(weatherDescription.includes('Snow')||weatherDescription.includes('Frezzing'))
    return playlists[0].snowy;
    else if(weatherDescription.includes('Cloudy')||weatherCodes.includes('Fog'))
    return playlists[0].cloudy;
    else 
    return playlists[0].sunny;

}

//////////////////////////////////to get a playlist details via id ////////////////////////////////////
function getPlaylist(playListId,token){
    const options={
        url:`https://api.spotify.com/v1/playlists/${playListId}`,
        headers:{
            Authorization: `Bearer ${token}`
        }
    }
   
    return axios.get(options.url,{headers:options.headers})
    .then(response=>{
        response.data;
        //console.log(response.data);
        const items=response.data.tracks.items;
        console.log(items);
    }
    )
    .catch(error=>{
        console.log(error);
    }
    )


}


////////////////////////////////////We be calling the weather code/////////////////////////////////////
app.get('/playlist', async(req, res) => {
    const {location}=req.query;
    try{
        //takes the description from the weather function and pushes it here 
        const weatherDescription= await weather(location);
        //takes the weather description and returns the playlist url
        const playLists= playList(weatherDescription)
         
        //play the song playlist in this case return the plsytlist url
        const playListId=await getPlaylist(playLists.split('/')[4]);
        //////////////////////////////////////////////////////////
        res.json({playListId});
        console.log(playListId);

    }
    catch(error){
        console.log(error);
        res.status(500).send('Some error occured');
    }
})




////////////////////////////////////////////app listening////////////////////////////////////////////
app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`);
    }
);


/*improvements to be made
1. add a location variable----- done
2. add a time variable
3. get real time weather data-----------done
4. take the playlist and list the songs in the playlist ---------- done
*/