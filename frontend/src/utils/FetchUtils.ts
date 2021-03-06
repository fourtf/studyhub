export const fetchJson = async (url: string, requestOptions: RequestInit, object2Send: object): Promise<any> => {
    requestOptions.headers = {...requestOptions.headers, 'Content-Type': 'application/json'}
    requestOptions = {...requestOptions, body: JSON.stringify(object2Send)}
    const response = await fetch(url, requestOptions)
    return response.json()
}

export const fetchJsonWithAuth = async (url: string, requestOptions: RequestInit, object2Send: object) : Promise<any> => {
    requestOptions.headers = {...requestOptions.headers, 'Content-Type': 'application/json' ,'Token': getCookie("studyhub_token")}
    requestOptions = {...requestOptions, body: JSON.stringify(object2Send)}
    const response = await fetch(url, requestOptions)
    return response.json()
}

const getCookie = (cookieName: string): string  => {
    let name = cookieName + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let ca = decodedCookie.split(';');
    for(let i = 0; i < ca.length; i++) {
      let c = ca[i];
      while (c.charAt(0) === ' ') {
        c = c.substring(1);
      }
      if (c.indexOf(name) === 0) {
        return c.substring(name.length, c.length);
      }
    }
    return "";
}