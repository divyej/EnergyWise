const express = require('express');
const axios = require('axios');
const Spotify= require('spotify-web-api-node');
const {SPOTIFY_CLIENT_ID, SPOTIFY_CLIENT_SECRET,API} = require('./config.js');


require('dotenv').config()

const app = express()
const port = 5000

const api=API

/////////////////////////////spotify api/////////////////////////////////////

const spotifyApi = new Spotify({
  clientId: SPOTIFY_CLIENT_ID,
  clientSecret: SPOTIFY_CLIENT_SECRET
});

spotifyApi
    .clientCredentialsGrant()
    .then(function(result) {
        console.log('It worked! Your access token is: ' + result.body.access_token)
    }).catch(function(err) {
        console.log("Something went wrong when retrieving an access token",err)
    });

/////////////////////////////tomoorrow api/////////////////////////////////////



app.use(express.json());

app.get('/weather', async(req, res) => {
    try{
        const response = await axios.get(`https://api.tomorrow.io/v4/timelines?location=40.75872069597532,-73.98529171943665&fields=temperature,precipitationIntensity&apikey=${api}`);
    
    const weatherData = response.data.data.timelines[0].intervals[0].values;
    const temperature = weatherData.temperature;

    console.log(temperature)
    console.log(weatherData);
    res.json(weatherData);
    }
    catch(error){
        console.log(error);
        res.status(500).send('Some error occured');
    }
})


app.get('/', (req, res) => {
    res.send('Hello World!');
    }
);

app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`);
    }
);