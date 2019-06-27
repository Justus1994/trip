const fetch = window.fetch;
 window.fetch = (...args) => (async(args) => {
     //intercept to add headers and url
     args[0] = 'api/' + args[0];
     let test = args[1];
       const opt = {
         method : typeof(test) === 'string' ? test : test.method,
         headers :  {'Authorization' : window.localStorage.getItem('Authorization-Token')}
     }
     args[1] = opt;
     var result = await fetch(...args);
     if(!result.ok){
       throw new Error('Something went wrong');
       return;
     }
       console.log(result.clone());
       let json = await result.json();
       return json;
 })(args);
