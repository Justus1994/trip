const fetch = window.fetch;
 window.fetch = (...args) => (async(args) => {
     //intercept req to add headers and change url
     args[0] = 'api/' + args[0];
     let oldOpt = args[1];
       const newOpt = {
         method : typeof(oldOpt) === 'string' ? oldOpt : oldOpt.method,
         headers :  {'Authorization' : window.localStorage.getItem('Authorization-Token')}
     }
     args[1] = newOpt;
     var result = await fetch(...args);
     //intercept resp throw err if no ok else return body as json
     if(!result.ok){
       throw new Error('Something went wrong');
       return;
     }
       let json = await result.json();
       return json;
 })(args);

export default fetch;
