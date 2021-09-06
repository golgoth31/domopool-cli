import axios from 'axios';

let boxURL: string;

// if (window.DOMOPOOL_HOST !== '{{ .DomopoolBoxHost }}') {
//     boxURL = window.DOMOPOOL_SCHEME + '://' + window.DOMOPOOL_HOST + ':' + window.DOMOPOOL_PORT;
// }
// else {
boxURL = window.location.protocol + '//' + window.location.host;
// }

const instance = axios.create({
    baseURL: boxURL
});

// Alter defaults after instance has been created
instance.defaults.headers.common['Access-Control-Allow-Origin'] = "*";

export default instance;
