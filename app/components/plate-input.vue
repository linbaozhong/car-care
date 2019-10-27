<template>
	<view class="plate-input">
		<view class="plate-input-header">
			<button type="primary" size="mini" plain="true" @click="onClose">关闭</button>
		</view>
		<view class="plate-input-body">
			<view class="plate-input-body-1" v-show="plate.length===0">
				<view class="plate-input-body-row">
					<button v-for="word in first" :key="word" size="mini" plain="true" @click="onWordClick(word)">{{word}}</button>
				</view>
				<view class="plate-input-body-row">
					<button v-for="word in second" :key="word" size="mini" plain="true" @click="onWordClick(word)">{{word}}</button>
				</view>
				<view class="plate-input-body-row">
					<button v-for="word in third" :key="word" size="mini" plain="true" @click="onWordClick(word)">{{word}}</button>
				</view>
				<view class="plate-input-body-row">
					<button v-for="word in forth" :key="word" size="mini" plain="true" @click="onWordClick(word)">{{word}}</button>
				</view>
			</view>
			<view class="plate-input-body-2" v-show="plate.length>0">
				<view class="plate-input-body-row">
					<button v-for="word in digit" :key="word" size="mini" plain="true" @click="onWordClick(word)">{{word}}</button>
				</view>
				<view class="plate-input-body-row">
					<button v-for="word in alpha1" :key="word" size="mini" plain="true" @click="onWordClick(word)">{{word}}</button>
				</view>
				<view class="plate-input-body-row">
					<button v-for="word in alpha2" :key="word" size="mini" plain="true" @click="onWordClick(word)">{{word}}</button>
				</view>
				<view class="plate-input-body-row">
					<button v-for="word in alpha3" :key="word" size="mini" plain="true" @click="onWordClick(word)">{{word}}</button>
				</view>
			</view>
			<button class="btn-delete iconfont icon-tuige" size="mini" plain="true" @click="onDeleteClick()"></button>
		</view>
	</view>
</template>

<script>
	export default {
		name: "plateinput",
		props:{
			plateNumber: {
					type: String,
					default: ''
				}
		},
		data() {
			return {
				hide:this.hidden,
				plate:this.plateNumber,
				first: ['京', '沪', '粤', '津', '冀', '晋', '蒙', '辽', '吉', '黑'],
				second: ['苏', '浙', '皖', '闽', '鲁', '豫', '鄂', '湘'],
				third: ['桂', '琼', '渝', '川', '贵', '云', '藏'],
				forth: ['陕', '甘', '青', '宁', '新'],
				digit: ['1', '2', '3', '4', '5', '6', '7', '8', '9', '0'],
				alpha1: ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K'],
				alpha2: ['L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'U', 'V'],
				alpha3: ['W', 'X', 'Y', 'Z', '港', '澳', '学'],
			};
		},
		methods:{
			onClose(e){
				this.hide = true
				this.$emit('on-close', e)
			},
			onWordClick(w){
				this.plate+=w
				this.$emit('on-change',this.plate)
			},
			onDeleteClick(){
				this.plate = this.plate.slice(0,this.plate.length-1)
				this.$emit('on-change',this.plate)
			}
		}
	}
</script>

<style lang="less">
	.plate-input {
		position: fixed;
		bottom: 0;
		left: 0;
		right: 0;
		height: 430rpx;
		border-top:1px solid #C8C7CC;
		background: #D1D5D9;

		.plate-input-header {
			height: 70rpx;
			background: #F7F7F7;
			display: flex;
			justify-content:flex-end;
			align-items:center;
			overflow:hidden;
			button{
				border: none;
				padding: 0;
				margin: 0 20rpx;
			}
		}

		.plate-input-body {
			position: relative;
			padding: 10rpx;
			overflow:hidden;
			.plate-input-body-1 .plate-input-body-row{
				justify-content: center;
			}
			.plate-input-body-row {
				margin: 10rpx 0;
				display: flex;

				button {
					width:66rpx ;
					height: 77rpx;
					line-height: 77rpx;
					background: #F7F7F7;
					padding: 0;
					margin: 0 6rpx;
					border-color: #999999;
				}
			}

			.btn-delete {
				position: absolute;
				height: 75rpx;
				line-height: 75rpx;
				bottom: 20rpx;
				right: 15rpx;
				padding: 0 15rpx;
				background: #F7F7F7;
				border-color: #999999;
			}
			.btn-delete.iconfont{
				font-size: 60rpx;
			}
		}
	}
</style>
