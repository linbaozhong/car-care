const baseUrl = {
	host: '127.0.0.1',
	port: 8080,
	protocol: 'http:',
	appId: 3
}

const SESSION_KEY = 'SESSION_USER'
const APP_DATA = getApp().globalData
APP_DATA.session = null
//
let session = {}

const getUrl = (path, auto) => {
	session = APP_DATA.session || session
	return baseUrl.protocol + '//' + baseUrl.host + ':' + baseUrl.port + path + '?tk=' + (
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
			header: {
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

const login = (callback = null) => {
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
						setSession(res.data, callback)
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
			APP_DATA.session = session
			if (callback) {
				callback(APP_DATA.session)
			}
		},
		fail(err) {
			// console.log(err)
		}
	})
}

const delSession = () => {
	uni.removeStorage({
		key: SESSION_KEY
	})
}

const initSession = () => {
	if (APP_DATA.session != null) {
		return
	}
	APP_DATA.session = getSession()
	console.log(APP_DATA.session)
	if (APP_DATA.session != null) {
		return
	}
	uni.reLaunch({
		url: '../../login/login'
	})
}

module.exports = {
	getSession,
	login,
	setSession,
	delSession,
	initSession,
	ajax,
	APP_DATA,
}
