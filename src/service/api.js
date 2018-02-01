import axios from 'axios'

const Http = {

    // 全権取得
    func(data) {
        return axios.post(`https://hal-iot.net/api/function`,data)
    }
}

export default Http;