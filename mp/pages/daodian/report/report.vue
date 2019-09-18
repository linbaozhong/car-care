<template>
	<view id="report">
		<view class="user_sign_url" @click="onShowSignature">
			<view v-if="report.user_sign_url==''">
				点我签名
			</view>
			<image v-else :src="report.user_sign_url" mode="aspectFit"></image>
		</view>
		<!-- 底部分享弹窗 -->
		<uni-popup ref="signature" type="bottom" :custom="true">
			<view class="popup-body">
				<view class="tools">
					<view class="" @click="_resetCanvas()">重写</view>
					<view class="" @click="onOk">确定</view>
				</view>
				<view>
					<canvas style="width: 750rpx; height: 250rpx;" canvas-id="firstCanvas" class="signature" @touchstart="onTouchstart"
					 @touchmove="onTouchmove" @touchend="onTouchend"></canvas>
				</view>
			</view>
		</uni-popup>
	</view>
</template>

<script>
	import uniPopup from '@/components/uni-popup/uni-popup.vue'
	let signature = {
		ctx: null,
		drawing: false,
		lastPos: {
			x: 0,
			y: 0
		},
		currentPos: {
			x: 0,
			y: 0
		},
	}
	export default {
		components: {
			uniPopup,
		},
		data() {
			return {
				type: 0,
				car: {
					id: 1,
					plate: '京N9CD10',
					brand: {
						name: '昊锐',
						logo_url: 'https://car3.autoimg.cn/cardfs/series/g29/M0B/AE/D4/100x100_f40_autohomecar__wKgHG1s9t3CAQfVMAABHgh0VMdw516.png',
						maker: '上海大众',
					},
					model: '手自一体 优雅版',
					years: '2012',
					odometer: 5600,
				},
				report: {
					id: 1,
					rpt_time: 1568780293449,
					abnormal: 2,
					clerk_sign_url: '',
					user_sign_url: '',
				},
			}
		},
		onLoad(opt) {
			this._resetCanvas()
			this.type = opt.type
			this.report.id = opt.id
		},
		methods: {
			onShowSignature() {
				this.$refs['signature'].open()
			},
			onTouchstart(e) {
				signature.drawing = true;
				signature.lastPos = this._getPosition(e)
				signature.currentPos = this._getPosition(e)
				this._renderCanvas()
			},
			onTouchmove(e) {
				signature.currentPos = this._getPosition(e)
				this._renderCanvas()
			},
			onTouchend(e) {
				signature.drawing = false
				// // Trigger a change event
				// var changedEvent = $.Event('jq.signature.changed');
				// this.$element.trigger(changedEvent);
			},
			onOk() {
				this._getDataURL()
				this.$refs['signature'].close()
			},
			_getPosition(e) {
				let pos = e.changedTouches[0]
				return {
					x: pos.x,
					y: pos.y,
				}
			},
			// Render the signature to the canvas
			_renderCanvas() {
				if (signature.drawing) {
					signature.ctx.moveTo(signature.lastPos.x, signature.lastPos.y)
					signature.ctx.lineTo(signature.currentPos.x, signature.currentPos.y)
					signature.ctx.stroke()
					signature.ctx.draw(true)
					signature.lastPos = signature.currentPos;
				}
			},
			// Reset the canvas context
			_resetCanvas() {
				signature.ctx = uni.createCanvasContext('firstCanvas')
				signature.ctx.setStrokeStyle("#ff0000")
				signature.ctx.setLineWidth(2)
				signature.ctx.draw()
			},
			_getDataURL() {
				let _this = this
				uni.canvasToTempFilePath({
					canvasId: 'firstCanvas',
					success(res) {
						if (res.errMsg === 'canvasToTempFilePath:ok') {
							_this.report.user_sign_url = res.tempFilePath
						} else {
							console.log(res.errMsg)
						}
					}
				})
			}
		}
	}
</script>

<style lang="scss">
	#report{
		.user_sign_url{
			width: 750upx;
			height: 300upx;
			display: flex;
			flex-direction: column;
			justify-content: space-between;
			align-items: center;
			border: 1px solid #dddddd;
		}
	}
	.popup-body {
		background: #FFFFFF;

		.tools {
			padding: 10upx 20upx;
			font-size: 32upx;
			display: flex;
			justify-content: space-between;
			align-items: center;
		}
	}

	.signature {
		box-sizing: border-box;
		border: 1px dashed rgb(170, 170, 170);
		background: url("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAlgAAADICAYAAAA0n5+2AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyZpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTM4IDc5LjE1OTgyNCwgMjAxNi8wOS8xNC0wMTowOTowMSAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTcgKFdpbmRvd3MpIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOjNGQjExNjZFNUMzRDExRThBMEVBODlEOEI3QUZEQTA3IiB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOjNGQjExNjZGNUMzRDExRThBMEVBODlEOEI3QUZEQTA3Ij4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6M0ZCMTE2NkM1QzNEMTFFOEEwRUE4OUQ4QjdBRkRBMDciIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6M0ZCMTE2NkQ1QzNEMTFFOEEwRUE4OUQ4QjdBRkRBMDciLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz4uosxwAAAJuUlEQVR42uza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") 0% 0% / cover transparent;
		cursor: crosshair;
	}
</style>
