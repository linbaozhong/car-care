<template>
	<view class="content">
		<view class="item">
			<view class="input-plate" @click="onPlateClick">
				<view class="placeholder">{{placeholder}}</view>
				<view class="input-plate-number">
					<view v-for="word in record.plateNumber" :key="word" class="input-plate-number-word">{{word}}</view>
				</view>
			</view>
		</view>
		<view class="item" @click="onVinChange(record.vin_url==='')">
			<view class="">
				车架号：
			</view>
			<image class="vin_url" v-if="record.vin_url!=''" :src="record.vin_url" mode="aspectFit" @click="onVinChange(true)"></image>
		</view>
		<view class="item" @click="onUserSignChange(record.user_sign_url==='')">
			<view class="">
				签名：
			</view>
			<image class="sign_url" v-if="record.user_sign_url!=''" :src="record.user_sign_url" mode="aspectFit" @click="onUserSignChange(true)"></image>
		</view>
		<!-- 底部车牌号输入弹窗 -->
		<plate_input v-if="showPlateInput" :plateNumber="record.plateNumber" @on-close="onPlateClose" @on-change="onPlateChange"></plate_input>
		<!-- 底部签名弹窗 -->
		<uni-popup ref="signature" type="bottom" :custom="true">
			<view class="popup-body signature">
				<view class="tools">
					<view class="" @click="_resetCanvas()">重写</view>
					<view class="palette">
						<view class="palette-color" :style="{background:color}"></view>
						<view v-for="color in colors" :key="color" :style="{background:color}" @click="onColorChange(color)"></view>
					</view>
					<view class="" @click="onOk">确定</view>
				</view>
				<view>
					<canvas style="width: 750rpx; height: 250rpx;" canvas-id="firstCanvas" class="signature-body" @touchstart="onTouchstart"
					 @touchmove="onTouchmove" @touchend="onTouchend"></canvas>
				</view>
				<view class="tools">
					
				</view>
			</view>
		</uni-popup>
	</view>
</template>

<script>
	import plate_input from '@/components/plate-input.vue'
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
			plate_input,
			uniPopup,
		},
		data() {
			return {
				colors:['#ff0000','#333333','#666666','#999999'],
				color:'#333333',
				placeholder: '车牌号：',
				showPlateInput: false,
				record: {
					plateNumber: '',
					vin_url: '',
					clerk_sign_url: '',
					user_sign_url: '',
				},
			}
		},
		watch: {
			"record.plateNumber"(v) {
				if (v === '') {
					this.placeholder = '车牌号：'
				} else {
					this.placeholder = ''
				}
			},
		},
		onLoad() {
			this._resetCanvas()
		},
		methods: {
			onPlateClick() {
				this.showPlateInput = true
			},
			onPlateClose() {
				this.showPlateInput = false
			},
			onPlateChange(v) {
				this.record.plateNumber = v
			},
			onVinChange(t) {
				this.onPlateClose()
				if (t) {
					let _this = this
					uni.chooseImage({
						count: 6, //默认9
						sizeType: ['original', 'compressed'], //可以指定是原图还是压缩图，默认二者都有
						success: function(res) {
							_this.record.vin_url = res.tempFilePaths[0];
						}
					});
				}
			},
			onUserSignChange(t) {
				this.onPlateClose()
				if (t) {
					this.$refs['signature'].open()
					this._resetCanvas()
				}
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
				signature.ctx.setStrokeStyle(this.color)
				signature.ctx.setLineWidth(3)
				signature.ctx.save()
			},
			_getDataURL() {
				let _this = this
				uni.canvasToTempFilePath({
					canvasId: 'firstCanvas',
					success(res) {
						if (res.errMsg === 'canvasToTempFilePath:ok') {
							_this.record.user_sign_url = res.tempFilePath
						} else {
							console.log(res.errMsg)
						}
					}
				})
			},
			onColorChange(color){
				this.color = color
				signature.ctx.setStrokeStyle(this.color)
			}
		}
	}
</script>

<style lang="scss">
	.content {
		padding: 20upx;

		.item {
			padding: 20upx 10upx;
			font-size: 32upx;
			min-height: 50upx;
			line-height: 50upx;
			border-bottom: 1px solid $uni-border-color;
			.sign_url{
				height: 180upx;
				max-height: 110px;
			}
			.vin_url{
				height: 200upx;
				max-height: 200px;
			}
		}

		.input-plate {
			position: relative;
			text-align: center;

			.placeholder {
				text-align: left;
			}

			.input-plate-number {
				position: absolute;
				top: 0;
				right: 50upx;
				bottom: 0;
				left: 50upx;
				display: flex;
				justify-content: center;

				.input-plate-number-word {
					width: 90upx;
					flex-shrink: 1;
				}
			}
		}
	}
	.signature {
		background: #FFFFFF;

		.tools {
			padding: 10upx 20upx;
			font-size: 32upx;
			display: flex;
			justify-content: space-between;
			align-items: center;
			.palette{
				display: flex;
				view{
					width: 40upx;
					height: 40upx;
				}
				.palette-color{
					margin-right: 20upx;
				}
			}
		}

		.signature-body {
			box-sizing: border-box;
			border: 1px dashed rgb(170, 170, 170);
			background: url("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAlgAAADICAYAAAA0n5+2AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyZpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTM4IDc5LjE1OTgyNCwgMjAxNi8wOS8xNC0wMTowOTowMSAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTcgKFdpbmRvd3MpIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOjNGQjExNjZFNUMzRDExRThBMEVBODlEOEI3QUZEQTA3IiB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOjNGQjExNjZGNUMzRDExRThBMEVBODlEOEI3QUZEQTA3Ij4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6M0ZCMTE2NkM1QzNEMTFFOEEwRUE4OUQ4QjdBRkRBMDciIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6M0ZCMTE2NkQ1QzNEMTFFOEEwRUE4OUQ4QjdBRkRBMDciLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz4uosxwAAAJuUlEQVR42uza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") 0% 0% / cover transparent;
			cursor: crosshair;
		}
	}
</style>
