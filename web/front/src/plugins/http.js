import Vue from 'vue'
import axios from 'axios'

// axios请求地址
axios.defaults.baseURL = 'http://42.192.21.110/api/v1'

Vue.prototype.$http = axios
