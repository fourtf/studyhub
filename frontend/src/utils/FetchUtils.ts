export const fetchJson = async <T>(url: string, requestOptions: RequestInit, object2Send: object): Promise<T> => {
    requestOptions.headers = {...requestOptions.headers, 'Content-Type': 'application/json'}
    requestOptions = {...requestOptions, body: JSON.stringify(object2Send)}
    const response = await fetch(url, requestOptions)
    return response.json() as Promise<T>
}

export const fetchJsonWithAuth = async <T>(url: string, requestOptions: RequestInit, object2Send: object) : Promise<T> => {
    requestOptions.headers = {...requestOptions.headers, 'Token': getCookie("studyhub_token")}
    return fetchJson(url, requestOptions, object2Send)
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