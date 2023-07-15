const express = require('express');
const axios = require('axios');

require('dotenv').config();

const app = express();
const port = 3000;

const api=process.env.API;

app.use(express.json());

app.get('/weather', async(req, res) => {
    try{
        const response = await axios.get(`https://api.tomorrow.io/v4/timelines?location=40.75872069597532,-73.98529171943665&fields=temperature,precipitationIntensity&apikey=${api}`);
    
    const weatherData = response.data.data.timelines[0].intervals[0].values;
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