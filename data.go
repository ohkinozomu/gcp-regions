package regions

const REGION_DATA string = `
{
	"asia-east1":{
		"name":"Taiwan",
		"flag": "https://upload.wikimedia.org/wikipedia/commons/7/72/Flag_of_the_Republic_of_China.svg",
		"latitude": 23.69781,
		"longitude": 120.960515
	},
	"asia-east2":{
		"name":"Hong Kong",
		"flag": "https://upload.wikimedia.org/wikipedia/commons/5/5b/Flag_of_Hong_Kong.svg",
		"latitude": 22.396428,
		"longitude": 114.109497
	},
	"asia-northeast1":{
		"name":"Tokyo",
		"flag": "https://upload.wikimedia.org/wikipedia/en/9/9e/Flag_of_Japan.svg",
		"latitude": 35.6762,
		"longitude": 139.6503
	},
	"asia-northeast2":{
		"name":"Osaka",
		"flag": "https://upload.wikimedia.org/wikipedia/en/9/9e/Flag_of_Japan.svg",
		"latitude": 34.6937,
		"longitude": 135.5023
	},
	"asia-northeast3":{
		"name":"Seoul",
		"flag": "https://upload.wikimedia.org/wikipedia/commons/0/09/Flag_of_South_Korea.svg",
		"latitude": 37.5665,
		"longitude": 126.9780
	},
	"asia-south1":{
		"name":"Delhi",
		"flag": "https://upload.wikimedia.org/wikipedia/en/4/41/Flag_of_India.svg",
		"latitude": 19.0760,
		"longitude": 72.8777
	},
	"asia-south2":{
		"name":"Mumbai",
		"flag": "https://upload.wikimedia.org/wikipedia/en/4/41/Flag_of_India.svg",
		"latitude": 28.7041,
		"longitude": 77.1025
	},
	"asia-southeast1":{
		"name":"Singapore",
		"flag": "https://upload.wikimedia.org/wikipedia/commons/4/48/Flag_of_Singapore.svg",
		"latitude": 1.3521,
		"longitude": 103.8198
	},
	"asia-southeast2":{
		"name":"Jakarta",
		"flag": "https://upload.wikimedia.org/wikipedia/commons/9/9f/Flag_of_Indonesia.svg",
		"latitude": -6.2088,
		"longitude": 106.8456
	},
	"australia-southeast1":{
		"name":"Sydney",
		"flag": "https://upload.wikimedia.org/wikipedia/commons/8/88/Flag_of_Australia_%28converted%29.svg",
		"latitude": -33.8688,
		"longitude": 151.2093
	},
	"australia-southeast2":{
		"name":"Melbourne",
		"flag": "https://upload.wikimedia.org/wikipedia/commons/8/88/Flag_of_Australia_%28converted%29.svg",
		"latitude": -37.8136,
		"longitude": 144.9631
	},
	"europe-central2":{
		"name":"Warsaw, Poland",
		"flag": "https://upload.wikimedia.org/wikipedia/en/1/12/Flag_of_Poland.svg",
		"latitude": 52.2297,
		"longitude": 21.0122
	},
	"europe-north1":{
		"name":"Hamina, Finland",
		"flag": "https://upload.wikimedia.org/wikipedia/commons/b/bc/Flag_of_Finland.svg",
		"latitude": 60.5693,
		"longitude": 27.1878
	},
	"europe-west1":{
		"name":"Belgium",
		"flag": "https://upload.wikimedia.org/wikipedia/commons/9/92/Flag_of_Belgium_%28civil%29.svg",
		"latitude": 50.5039,
		"longitude": 4.4699
	},
	"europe-west2":{
		"name":"London, UK",
		"flag": "https://upload.wikimedia.org/wikipedia/en/a/ae/Flag_of_the_United_Kingdom.svg",
		"latitude": 51.5074,
		"longitude": -0.1278
	},
	"europe-west3":{
		"name":"Frankfurt, Germany",
		"flag": "https://upload.wikimedia.org/wikipedia/en/b/ba/Flag_of_Germany.svg",
		"latitude": 50.1109,
		"longitude": 8.6821
	},
	"europe-west4":{
		"name":"Netherlands",
		"flag": "https://upload.wikimedia.org/wikipedia/commons/2/20/Flag_of_the_Netherlands.svg",
		"latitude": 52.1326,
		"longitude": 5.2913
	},
	"europe-west6":{
		"name":"Zurich, Switzerland",
		"flag": "https://upload.wikimedia.org/wikipedia/commons/f/f3/Flag_of_Switzerland.svg",
		"latitude": 47.3769,
		"longitude": 8.5417
	},
	"northamerica-northeast1":{
		"name":"Montréal, Canada",
		"flag": "https://upload.wikimedia.org/wikipedia/commons/d/d9/Flag_of_Canada_%28Pantone%29.svg",
		"latitude": 45.5017,
		"longitude": -73.5673
	},
	"southamerica-east1":{
		"name":"São Paulo",
		"flag": "https://upload.wikimedia.org/wikipedia/en/0/05/Flag_of_Brazil.svg",
		"latitude": -23.5505,
		"longitude": -46.6333
	},
	"us-central1":{
		"name":"Iowa, USA",
		"flag": "https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg",
		"latitude": 41.8780,
		"longitude": -93.0977
	},
	"us-east1":{
		"name":"South Carolina, USA",
		"flag": "https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg",
		"latitude": 33.8361,
		"longitude": -81.1637
	},
	"us-east4":{
		"name":"Northern Virginia, USA",
		"flag": "https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg",
		"latitude": 38.8334,
		"longitude": -77.2365
	},
	"us-west1":{
		"name":"Oregon, USA",
		"flag": "https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg",
		"latitude": 43.8041,
		"longitude": -120.5542
	},
	"us-west2":{
		"name":"Los Angeles, USA",
		"flag": "https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg",
		"latitude": 34.0522,
		"longitude": -118.2437
	},
	"us-west3":{
		"name":"Salt Lake City, USA",
		"flag": "https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg",
		"latitude": 40.7608,
		"longitude": -111.8910
	},
	"us-west4":{
		"name":"Las Vegas, USA",
		"flag": "https://upload.wikimedia.org/wikipedia/en/a/a4/Flag_of_the_United_States.svg",
		"latitude": 36.1699,
		"longitude": -115.1398
	}
}`
