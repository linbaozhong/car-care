(global["webpackJsonp"]=global["webpackJsonp"]||[]).push([["pages/tabbar-3-detial/tabbar-3-meirong/tabbar-3-meirong"],{"137e":function(t,n,e){"use strict";(function(t,o){Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0;var a=function(){return e.e("components/plate-input").then(e.bind(null,"9369"))},r=function(){return e.e("components/uni-popup/uni-popup").then(e.bind(null,"e511"))},s={ctx:null,drawing:!1,lastPos:{x:0,y:0},currentPos:{x:0,y:0}},i={components:{plate_input:a,uniPopup:r},data:function(){return{colors:["#ff0000","#333333","#666666","#999999"],color:"#333333",placeholder:"车牌号：",showPlateInput:!1,record:{plateNumber:"",vin_url:"",clerk_sign_url:"",user_sign_url:""}}},watch:{"record.plateNumber":function(t){this.placeholder=""===t?"车牌号：":""}},onLoad:function(){this._resetCanvas()},methods:{onPlateClick:function(){this.showPlateInput=!0},onPlateClose:function(){this.showPlateInput=!1},onPlateChange:function(t){this.record.plateNumber=t},onVinChange:function(n){if(this.onPlateClose(),n){var e=this;t.chooseImage({count:6,sizeType:["original","compressed"],success:function(t){e.record.vin_url=t.tempFilePaths[0]}})}},onUserSignChange:function(t){this.onPlateClose(),t&&this.$refs["signature"].open()},onTouchstart:function(t){s.drawing=!0,s.lastPos=this._getPosition(t),s.currentPos=this._getPosition(t),this._renderCanvas()},onTouchmove:function(t){s.currentPos=this._getPosition(t),this._renderCanvas()},onTouchend:function(t){s.drawing=!1},onOk:function(){this._getDataURL(),this.$refs["signature"].close()},_getPosition:function(t){var n=t.changedTouches[0];return{x:n.x,y:n.y}},_renderCanvas:function(){s.drawing&&(s.ctx.moveTo(s.lastPos.x,s.lastPos.y),s.ctx.lineTo(s.currentPos.x,s.currentPos.y),s.ctx.stroke(),s.ctx.draw(!0),s.lastPos=s.currentPos)},_resetCanvas:function(){s.ctx=t.createCanvasContext("firstCanvas"),s.ctx.setStrokeStyle(this.color),s.ctx.setLineWidth(3),s.ctx.draw()},_getDataURL:function(){var n=this;t.canvasToTempFilePath({canvasId:"firstCanvas",success:function(t){"canvasToTempFilePath:ok"===t.errMsg?n.record.user_sign_url=t.tempFilePath:console.log(o(t.errMsg," at pages\\tabbar-3-detial\\tabbar-3-meirong\\tabbar-3-meirong.vue:176"))}})},onColorChange:function(t){this.color=t,s.ctx.setStrokeStyle(this.color)}}};n.default=i}).call(this,e("6e42")["default"],e("0de9")["default"])},"236e":function(t,n,e){"use strict";e.r(n);var o=e("137e"),a=e.n(o);for(var r in o)"default"!==r&&function(t){e.d(n,t,function(){return o[t]})}(r);n["default"]=a.a},2984:function(t,n,e){},"8d65":function(t,n,e){"use strict";var o=e("2984"),a=e.n(o);a.a},a355:function(t,n,e){"use strict";e.r(n);var o=e("a45a"),a=e("236e");for(var r in a)"default"!==r&&function(t){e.d(n,t,function(){return a[t]})}(r);e("8d65");var s=e("2877"),i=Object(s["a"])(a["default"],o["a"],o["b"],!1,null,null,null);n["default"]=i.exports},a45a:function(t,n,e){"use strict";var o=function(){var t=this,n=t.$createElement;t._self._c},a=[];e.d(n,"a",function(){return o}),e.d(n,"b",function(){return a})},f12e:function(t,n,e){"use strict";(function(t){e("744b"),e("921b");o(e("66fd"));var n=o(e("a355"));function o(t){return t&&t.__esModule?t:{default:t}}t(n.default)}).call(this,e("6e42")["createPage"])}},[["f12e","common/runtime","common/vendor"]]]);