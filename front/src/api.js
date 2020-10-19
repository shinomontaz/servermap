
import axios from "axios";

const API_URL = 'http://'+process.env.REACT_APP_BACKEND_IP+':'+process.env.REACT_APP_BACKEND_PORT

class Api {
  apiCall = async ({ path, method = "get", data}) => {
    console.log(API_URL)
      const resp = await Promise.race([
        axios({
          method,
          data,
          url: API_URL + path
        }),
        new Promise((resolve, reject) => {
          setTimeout(
            () => reject(new Error("Ошибка подключения к интернету")), // TODO: сделать перевод по словарю в зависимости от языка
            30000
          );
        })
      ]);

console.log(resp.data)
      return resp.data;
  };

loadHosts = async ( vmtype ) => {
  var { data: hosts } = await axios.get(API_URL, {params: { vmtype: vmtype }} );
  return hosts
 }
    // this.apiCall({
    //   data: {
    //
    //   },
    // });
}

const api = new Api();
export default api;
