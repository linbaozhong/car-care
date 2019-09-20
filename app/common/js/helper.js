const baseUrl = {
	host: '127.0.0.1',
	port: 9000,
	protocol: 'http:',
	appId: 3
}

const SESSION_KEY = 'SESSION_USER'
//
let session = {
	token: '',
	auto: ''
}

const getUrl = (path, auto) => {
	session = getApp().globalData.session || session
	return baseUrl.protocol + '//' + baseUrl.host + ':' + baseUrl.port + path + '?appId=' + baseUrl.appId + '&tk=' + (
		auto ? (session.auto || '') : (session.token || ''))
}

const ajax = {
	get(url, data, callback) {
		return uni.request({
			url: getUrl(url),
			data: data,
			method: 'GET',
			success(res) {
				if (res.statusCode === 200) {
					callback(res.data)
					return
				}
				callback({
					success: false,
					error: {
						code: res.statusCode,
						detail: res.errMsg
					}
				})
			},
			fail(err) {
				callback({
					success: false,
					error: err
				})
			}
		});
	},
	post(url, data, callback) {
		return uni.request({
			url: getUrl(url),
			data: data,
			method: 'POST',
			header:{
				'content-type': 'application/x-www-form-urlencoded',
			},
			success(res) {
				if (res.statusCode === 200) {
					callback(res.data)
					return
				}
				callback({
					success: false,
					error: {
						code: res.statusCode,
						detail: res.errMsg
					}
				})
			},
			fail(err) {
				callback({
					success: false,
					error: err
				})
			}
		});
	}
}

const login = (callback=null) => {
	let provider = ''
	// #ifdef MP-WEIXIN
	provider = 'weixin'
	// #endif

	uni.login({
		provider: provider,
		success(res) {
			if (res.code) {
				ajax.get('/v1/weixin/login', {
					code: res.code,
				}, res => {
					if (res.success) {
						setSession(res.data,callback)
					}
				})
			}
		}
	})
}

const getSession = () => {
	try {
		return Object.assign({
			key: SESSION_KEY
		}, JSON.parse(uni.getStorageSync(SESSION_KEY)))
	} catch (e) {}
	return null
}
const setSession = (data, callback) => {
	if (session) {
		session = Object.assign({}, session, data)
	} else {
		session = data
	}
	uni.setStorage({
		key: SESSION_KEY,
		data: JSON.stringify(session),
		success(res) {
			getApp().globalData.session = session
			if (callback) {
				callback(getApp().globalData.session)
			}
		},
		fail(err) {
			// console.log(err)
		}
	})
}

module.exports = {
	getSession,
	login,
	setSession,
	ajax,
}
