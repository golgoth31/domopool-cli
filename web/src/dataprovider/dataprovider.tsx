import axios from 'axios';

let boxURL: string;
if (window.DOMOPOOL_IP != '{{ DomopoolBoxIP }}') {
  boxURL = window.DOMOPOOL_SCHEME + '://' + window.DOMOPOOL_IP + ':' + window.DOMOPOOL_IP
}
else {
  boxURL = window.location.protocol + '//' + window.location.host
}
export default axios.create({
  baseURL: boxURL
});
