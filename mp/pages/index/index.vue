<template>
	<view class="content">
		<!-- <signature :width="750" :height="250"></signature> -->
		<canvas style="width: 750rpx; height: 250rpx;" canvas-id="firstCanvas" class="signature" @touchstart="onTouchstart"
		@touchmove="onTouchmove"
		@touchend="onTouchend"></canvas>
	</view>
</template>

<script>
	let ctx
	export default {
		data() {
			return {
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
		},
		onLoad() {
			ctx = uni.createCanvasContext('firstCanvas')
			// this._resetCanvas()
			ctx.arc(100, 75, 50, 0, 2 * Math.PI)
			ctx.setFillStyle('#ff0000')
			ctx.fill()
			ctx.draw()
		},
		methods: {
			onTouchstart(e){
				// console.log(e)
				this.drawing = true;
				this.lastPos = this._getPosition(e)
				this.currentPos = this._getPosition(e)
				this._renderCanvas()
			},
			onTouchmove(e){
				// console.log(e)
				this.currentPos = this._getPosition(e)
				this._renderCanvas()
			},
			onTouchend(e){
				console.log(e)
				this.drawing = false
				// // Trigger a change event
				// var changedEvent = $.Event('jq.signature.changed');
				// this.$element.trigger(changedEvent);
			},
			_getPosition(e){
				let pos = e.changedTouches[0]
				return {
					x:pos.x,
					y:pos.y,
				}
			},
			// Render the signature to the canvas
			_renderCanvas() {
			  if (this.drawing) {
			    ctx.moveTo(this.lastPos.x, this.lastPos.y)
			    ctx.lineTo(this.currentPos.x, this.currentPos.y)
			    ctx.stroke()
				ctx.draw()
			    this.lastPos = this.currentPos;
			  }
			},
			// Reset the canvas context
			_resetCanvas() {
			  ctx = uni.createCanvasContext('firstCanvas')
			  ctx.setStrokeStyle("#ff0000")
			  ctx.setLineWidth(5)
			},
		}
	}
</script>

<style>
	.content {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.signature {
		margin-top: 100upx;
		box-sizing: border-box;
		border: 1px dashed rgb(170, 170, 170);
		background: url("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAlgAAADICAYAAAA0n5+2AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyZpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTM4IDc5LjE1OTgyNCwgMjAxNi8wOS8xNC0wMTowOTowMSAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTcgKFdpbmRvd3MpIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOjNGQjExNjZFNUMzRDExRThBMEVBODlEOEI3QUZEQTA3IiB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOjNGQjExNjZGNUMzRDExRThBMEVBODlEOEI3QUZEQTA3Ij4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6M0ZCMTE2NkM1QzNEMTFFOEEwRUE4OUQ4QjdBRkRBMDciIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6M0ZCMTE2NkQ1QzNEMTFFOEEwRUE4OUQ4QjdBRkRBMDciLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz4uosxwAAAJuUlEQVR42uza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") 0% 0% / cover transparent;
		cursor: crosshair;
	}
</style>
