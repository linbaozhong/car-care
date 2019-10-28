<script>
	export default {
		globalData:{
			session:null,
		},
		onLaunch: function() {
			// setTimeout(() => {
			// 	uni.setTabBarBadge({
			// 		index: 1,
			// 		text: '3'
			// 	});
			// 	uni.showTabBarRedDot({
			// 		index: 3
			// 	});
			// }, 1000);
		},
		onShow: function() {
			
		},
		onHide: function() {
		}
	};
</script>

<style>
	/*每个页面公共css */
	/* uni-app默认全局使用flex布局。因为flex布局有利于跨更多平台，尤其是采用原生渲染的平台。如不了解flex布局，请参考http://www.w3.org/TR/css3-flexbox/。如不使用flex布局，请删除或注释掉本行。*/

	@import "./common/css/reset.css";

	body,
	page {
		min-height: 100%;
		display: flex;
		font-size: 28rpx;
	}
	
	/* 原生组件模式下需要注意组件外部样式 */
	m-input {
		width: 100%;
		min-height: 100%;
		display: flex;
	}

	.content {
		display: flex;
		flex: 1;
		flex-direction: column;
		background-color: #efeff4;
		padding: 20rpx 0;
	}

	.input-group {
		background-color: #ffffff;
		margin: 0 20rpx;
		position: relative;
	}

	.input-group::before {
		position: absolute;
		right: 0;
		top: 0;
		left: 0;
		height: 1rpx;
		content: '';
		-webkit-transform: scaleY(.5);
		transform: scaleY(.5);
		background-color: #c8c7cc;
	}

	.input-group::after {
		position: absolute;
		right: 0;
		bottom: 0;
		left: 0;
		height: 1rpx;
		content: '';
		-webkit-transform: scaleY(.5);
		transform: scaleY(.5);
		background-color: #c8c7cc;
	}

	.input-row {
		display: flex;
		flex-direction: row;
		position: relative;
	}

	.input-row .title {
		width: 20%;
		height: 50rpx;
		min-height: 50rpx;
		padding: 15rpx 0;
		padding-left: 30rpx;
		line-height: 50rpx;
	}

	.input-row.border::after {
		position: absolute;
		right: 0;
		bottom: 0;
		left: 15rpx;
		height: 1rpx;
		content: '';
		-webkit-transform: scaleY(.5);
		transform: scaleY(.5);
		background-color: #c8c7cc;
	}

	.btn-row {
		margin-top: 50rpx;
		padding: 20rpx;
	}

	button.primary {
		background-color: #0faeff;
	}
</style>
