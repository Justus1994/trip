export default function prepareTripForSharing(trip){
    let shareString = "Schau dir meinen Trip an:\n" + trip.Nodes[0].location.country + "\n";
    let i = 1;
    trip.Nodes.forEach(node => {
        shareString += i++ + ". " + getTitle(node) + getCountry(node) + 
        "\nBild: " + node.urls.full + 
        "\nLocation: " + getLocationUrl(node) + "\n\n"
    });
    copyToClipboard(shareString);
};

function getTitle(node) {
    let title;
    if (node.location.title !== "" || node.location.title !== undefined) {
        title = node.location.title;
    } else {
        title = "Somewhere" 
    }
    return title;
}

function getCountry(node) {
    let country;
    node.location.country
    if (node.location.country === "" || node.location.country === undefined) {
        country = ", Somewhere";
    } else {
        country = node.location.country;
    }
    return country;
}

function getLocationUrl(node){
    let url;
    let lat = node.location.position.latitude;
    let long = node.location.position.longitude;
    if(lat === undefined && node.location.hasOwnProperty("city")){
      url = 'https://www.google.com/maps/search/?api=1&query=' + node.location.city;
    }
    else{
      url = 'https://www.google.com/maps/search/?api=1&query=' + node.location.country;
    }
    if(lat){
      url = 'http://maps.google.com/maps?t=k&q=loc:' + lat + '+' + long;
    }
    return url;
};

function copyToClipboard(shareString) {
    const el = document.createElement('textarea');
    el.value = shareString;
    document.body.appendChild(el);
    el.select();
    document.execCommand('copy');
    document.body.removeChild(el);
};