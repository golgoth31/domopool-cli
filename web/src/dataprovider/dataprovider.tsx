import axios from 'axios';

let boxURL: string;

if (window.DOMOPOOL_HOST != '{{ .DomopoolBoxHost }}') {
  boxURL = window.DOMOPOOL_SCHEME + '://' + window.DOMOPOOL_HOST + ':' + window.DOMOPOOL_PORT;
}
else {
  boxURL = window.location.protocol + '//' + window.location.host;
}
export default axios.create({
  baseURL: boxURL
});
