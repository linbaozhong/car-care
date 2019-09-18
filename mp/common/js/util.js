function formatTime(time) {
	if (typeof time !== 'number' || time < 0) {
		return time
	}

	var hour = parseInt(time / 3600)
	time = time % 3600
	var minute = parseInt(time / 60)
	time = time % 60
	var second = time

	return ([hour, minute, second]).map(function(n) {
		n = n.toString()
		return n[1] ? n : '0' + n
	}).join(':')
}

function formatLocation(longitude, latitude) {
	if (typeof longitude === 'string' && typeof latitude === 'string') {
		longitude = parseFloat(longitude)
		latitude = parseFloat(latitude)
	}

	longitude = longitude.toFixed(2)
	latitude = latitude.toFixed(2)

	return {
		longitude: longitude.toString().split('.'),
		latitude: latitude.toString().split('.')
	}
}

function getDuration(ms) {
	var days = ms / 1000 / 60 / 60 / 24;
	var daysRound = Math.floor(days);
	var hours = ms / 1000 / 60 / 60 - (24 * daysRound);
	var hoursRound = Math.floor(hours);
	var minutes = ms / 1000 / 60 - (24 * 60 * daysRound) - (60 * hoursRound);
	var minutesRound = Math.floor(minutes);
	var seconds = ms / 1000 - (24 * 60 * 60 * daysRound) - (60 * 60 * hoursRound) - (60 * minutesRound);
	var secondsRound = Math.floor(seconds)
	var time = (daysRound > 0 ? daysRound + '天 ' : '') + hoursRound + ':' + minutesRound + ':' + secondsRound
	return time;
}

var dateUtils = {
	UNITS: {
		'年': 31557600000,
		'月': 2629800000,
		'天': 86400000,
		'小时': 3600000,
		'分钟': 60000,
		'秒': 1000
	},
	humanize: function(milliseconds) {
		var humanize = '';
		for (var key in this.UNITS) {
			if (milliseconds >= this.UNITS[key]) {
				humanize = Math.floor(milliseconds / this.UNITS[key]) + key + '前';
				break;
			}
		}
		return humanize || '刚刚';
	},
	format: function(milliseconds) {
		var date = new Date(milliseconds)
		var diff = Date.now() - milliseconds;
		if (diff < this.UNITS['天']) {
			return this.humanize(diff);
		}
		var _format = function(number) {
			return (number < 10 ? ('0' + number) : number);
		};
		return date.getFullYear() + '/' + _format(date.getMonth() + 1) + '/' + _format(date.getDay()) + '-' +
			_format(date.getHours()) + ':' + _format(date.getMinutes());
	},
	parse: function(str) { //将"yyyy-mm-dd HH:MM:ss"格式的字符串，转化为一个Date对象
		var a = str.split(/[^0-9]/);
		return new Date(a[0], a[1] - 1, a[2], a[3], a[4], a[5]);
	}
};

function toNumber(data) {
	let val = ''
	if (typeof data === 'object' && data.detail) {
		val = data.detail.value
	} else {
		val = data
	}
	val = Number.parseInt(val)
	if (Number.isNaN(val)) return ''
	return val
}

function toPassword(d) {
	let p = []
	for (var i = 0; i < d.length; i++) {
		p.push('.')
	}
	return p.join('')
}

function isMobile(d) {
	return d.match(/^1[345678]\d{9}$/)
}

function isPassword(d) {
	return d.match(/^\d{4}$/)
}

module.exports = {
	isPassword,
	isMobile: isMobile,
	toNumber: toNumber,
	toPassword: toPassword,
	formatTime: formatTime,
	formatLocation: formatLocation,
	dateUtils: dateUtils,
	getDuration
}
