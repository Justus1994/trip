export default async function fetchData(url, method){

    url = 'api/' + url;
    let response = await fetch(url, {
                            method: method,
                            headers:{
                              'Authorization' : window.localStorage.getItem('Authorization-Token')
                            }
                          });
    if(response.status >= 300){
      return response.status;
    }
    let json = await response.json();
    return json;

}
