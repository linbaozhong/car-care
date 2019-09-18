import helper from '@/common/js/helper.js'

let accessid = ''
let policyBase64 = ''
let signature = ''
let callbackbody = ''
let key = ''
let host = ''
let expire = 0
// let g_object_name = ''

// 生成随机字符串
function random_string(len) {
	len = len || 32;
	var chars = 'ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678';
	var maxPos = chars.length;
	var pwd = '';
	for (let i = 0; i < len; i++) {
		pwd += chars.charAt(Math.floor(Math.random() * maxPos));
	}
	return pwd;
}
// 获取用户上传原始文件名
function get_file_name(filename) {
	let pos = filename.lastIndexOf('.')
	let suffix = ''
	if (pos != -1) {
		suffix = filename.substring(pos)
	}
	return suffix;
}
//  把随机生成的字符串拼接在原始上传文件名后面生成新的唯一文件名
function get_object_name(filename) {
	let suffix = get_file_name(filename)
	return key + random_string(10) + suffix
}

function upload(filename = [], callback = null) {

	let g_object_name = get_object_name(filename)
	// 上传
	uni.uploadFile({
		url: host,
		filePath: filename,
		name: 'file',
		formData: {
			'host': host,
			'key': g_object_name,
			'policy': policyBase64,
			'OSSAccessKeyId': accessid,
			'success_action_status': '200', //让服务端返回200,不然，默认会返回204
			//   'callback': callbackbody,
			'signature': signature,
		},
		success(res) {
			res = Object.assign({}, res, {
				url: host + '/' + g_object_name,
			})
			if (callback) {
				callback(res)
				return
			}
			console.log(res)
		}
	})
}

function getOSSSignAndUpload(filename, callback) {
	// 可以判断当前expire是否超过了当前时间， 如果超过了当前时间， 就重新取一下， 3 s 作为缓冲。
	if (expire < Date.parse(new Date()) / 1000 + 3) {
		// 调用后端服务器接口获取签名信息
		helper.ajax.get('/v1/oss/get', {}, res => {
			if (res.success) {
				let data = res.data
				policyBase64 = data.policy // 用户表单上传的策略（Policy)
				accessid = data.accessid // 用户请求的accessid
				signature = data.signature // 签名信息
				expire = parseInt(data.expire) // 上传策略Policy失效时间
				// callbackbody = data.callback // 回调
				host = data.host // 上传文件服务器地址
				key = data.dir // 上传文件的存储位置

				upload(filename, callback)
			} else {
				console.log(res.error)
			}
		})
	} else {
		upload(filename, callback)
	}
}
// 获取后端返回的签名信息，生成oss参数
function oss(files = [], callback = null) {
	if (!files || files.length===0) {
		return
	}
	if (typeof files === 'object') {
		files.map(filename => {
			getOSSSignAndUpload(filename,callback)
		})
	}else{
		getOSSSignAndUpload(files,callback)
	}
}

export {
	oss
}
