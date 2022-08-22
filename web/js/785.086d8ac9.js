"use strict";(self["webpackChunkfrontend"]=self["webpackChunkfrontend"]||[]).push([[785],{1208:(e,a,t)=>{t.d(a,{vk:()=>n,He:()=>r,pT:()=>s,Tb:()=>c,Be:()=>i});var o=t(1768);const n={cId:0,mId:0};function r(e,a){return new Promise(((t,n)=>{const r={type:"LOGIN_REQUEST",username:e,password:a};o.api.post("/mappa/login",r).then((e=>{t(e.data)})).catch((e=>{n(e)}))}))}function s(e,a,t){return new Promise(((n,r)=>{const s={cId:parseInt(`${e}`),mId:parseInt(`${a}`),msg:t};console.log("Sending auth",s),o.api.post("/tg/pub",s).then((e=>{console.log("SendAuthToChat",e),n(e)})).catch((e=>{console.error("SendAuthToChat",e),r(e)}))}))}function c(){return new Promise(((e,a)=>{o.api.get("/hc").then((a=>e(a.data))).catch((e=>{a(e)}))}))}function i(e){try{const a=atob(e),t=JSON.parse(a);if(!t.cId)throw new Error(`Invalid context data: ${JSON.stringify(t)}`);return{cId:t.cId,mId:t.mId}}catch(a){return console.error("ParseContent",a),n}}},4785:(e,a,t)=>{t.r(a),t.d(a,{default:()=>P});var o=t(3673);const n={class:"column"},r=(0,o._)("div",{class:"row"},[(0,o._)("h5",{class:"text-h5 text-white q-my-md"},"Autenticação mAPPa")],-1),s={class:"row"},c=(0,o._)("p",{class:"text-grey-16"},[(0,o.Uk)(" Sua senha não será armazenada."),(0,o._)("br"),(0,o.Uk)("Esta autenticação é usada para se obter acesso aos dados do mAPPa. ")],-1),i=(0,o._)("div",{class:"text-h6"},"Chave de autorização",-1),l=(0,o._)("p",{class:"text-grey-16"},[(0,o.Uk)(" Esta chave será utilizada para autorização no grupo da sua seção no Telegram."),(0,o._)("br"),(0,o.Uk)(" Copie e cole no grupo para que o nosso robô faça a conexão corretamente. ")],-1);function u(e,a,t,u,d,p){const m=(0,o.up)("q-input"),h=(0,o.up)("q-form"),f=(0,o.up)("q-card-section"),g=(0,o.up)("q-btn"),w=(0,o.up)("q-card-actions"),y=(0,o.up)("q-card"),q=(0,o.up)("q-separator"),v=(0,o.up)("q-page");return(0,o.wg)(),(0,o.j4)(v,{class:"window-height window-width row justify-center items-center"},{default:(0,o.w5)((()=>[(0,o._)("div",n,[r,(0,o._)("div",s,[e.authKey?((0,o.wg)(),(0,o.j4)(y,{key:1,square:"",bordered:"",class:"q-pa-lg shadow-1"},{default:(0,o.w5)((()=>[(0,o.Wm)(f,null,{default:(0,o.w5)((()=>[i])),_:1}),(0,o.Wm)(q,{inset:""}),(0,o.Wm)(f,null,{default:(0,o.w5)((()=>[(0,o.Wm)(m,{square:"",filled:"",clearable:"",modelValue:e.authKey,"onUpdate:modelValue":a[2]||(a[2]=a=>e.authKey=a),type:"text",label:"Chave",readonly:""},null,8,["modelValue"])])),_:1}),(0,o.Wm)(w,{class:"q-px-md"},{default:(0,o.w5)((()=>[(0,o.Wm)(g,{unelevated:"",color:"positive",size:"lg",class:"full-width",label:"Copiar",onClick:e.clickCopy,icon:"content_copy"},null,8,["onClick"])])),_:1}),(0,o.Wm)(f,{class:"text-center q-pa-none"},{default:(0,o.w5)((()=>[l])),_:1})])),_:1})):((0,o.wg)(),(0,o.j4)(y,{key:0,square:"",bordered:"",class:"q-pa-lg shadow-1"},{default:(0,o.w5)((()=>[(0,o.Wm)(f,null,{default:(0,o.w5)((()=>[(0,o.Wm)(h,{class:"q-gutter-md"},{default:(0,o.w5)((()=>[(0,o.Wm)(m,{square:"",filled:"",clearable:"",modelValue:e.username,"onUpdate:modelValue":a[0]||(a[0]=a=>e.username=a),type:"text",label:"Usuário mAPPa"},null,8,["modelValue"]),(0,o.Wm)(m,{square:"",filled:"",clearable:"",modelValue:e.password,"onUpdate:modelValue":a[1]||(a[1]=a=>e.password=a),type:"password",label:"Senha"},null,8,["modelValue"])])),_:1})])),_:1}),(0,o.Wm)(w,{class:"q-px-md"},{default:(0,o.w5)((()=>[(0,o.Wm)(g,{unelevated:"",color:"positive",size:"lg",class:"full-width",label:"Login",onClick:e.clickLogin},null,8,["onClick"])])),_:1}),(0,o.Wm)(f,{class:"text-center q-pa-none"},{default:(0,o.w5)((()=>[c])),_:1})])),_:1}))])])])),_:1})}var d=t(8825),p=t(1208),m=function(e,a,t,o){function n(e){return e instanceof t?e:new t((function(a){a(e)}))}return new(t||(t=Promise))((function(t,r){function s(e){try{i(o.next(e))}catch(a){r(a)}}function c(e){try{i(o["throw"](e))}catch(a){r(a)}}function i(e){e.done?t(e.value):n(e.value).then(s,c)}i((o=o.apply(e,a||[])).next())}))};const h=(0,o.aZ)({name:"MappaAuth",components:{},data(){return{username:"",password:"",authKey:"",context:p.vk,clipboardEnabled:!1,q:(0,d.Z)()}},mounted(){this.context=(0,p.Be)(this.$route.params.context)},methods:{clickCopy(){navigator.clipboard.writeText(this.authKey).then((()=>{this.q.notify({caption:"Área de transferência",message:"Chave foi copiada com sucesso. Você pode colá-la no grupo do telegram para continuar o setup.",icon:"check_circle",color:"success"})}),(()=>{this.q.notify({caption:"Área de transferência",message:"Erro na cópia para área de transferência. Faça o processo manualmente.",icon:"warning",color:"negative"})}))},clickLogin(){return m(this,void 0,void 0,(function*(){if(this.username&&this.password)try{this.q.notify({caption:"Conectando",message:"Login no proxy mAPPa",icon:"link",group:!1,color:"primary"});let e=yield(0,p.He)(this.username,this.password);if(!e)return void this.q.notify({caption:"Erro",message:"Falha na conexão com o proxy mAPPa!",icon:"warning",color:"negative"});let a={id:e.id,ttl:e.ttl,created:e.created,userId:e.userId,cId:this.context.cId,mId:this.context.mId},t=JSON.stringify(a);this.authKey=`/auth ${btoa(t)}`,(0,p.pT)(this.context.cId,this.context.mId,this.authKey).then((()=>{this.q.notify({caption:"Sucesso",message:"Autorização enviada para o chat no Telegram",icon:"ok",color:"success"})})).catch((e=>{this.q.notify({caption:"Erro",message:`Falha no envio da autorização para o chat no Telegram (${JSON.stringify(e)})!`,icon:"warning",color:"negative"})}))}catch(e){console.error("Login error:",e),this.q.notify({caption:"Erro",message:e.message,icon:"danger",group:!1})}else this.q.notify({caption:"Erro",message:"Informe usuário e senha para continuar!",icon:"warning",color:"negative"})}))}}});var f=t(4260),g=t(4379),w=t(151),y=t(5589),q=t(5269),v=t(1700),x=t(9367),b=t(3187),k=t(5869),I=t(7518),_=t.n(I);const C=(0,f.Z)(h,[["render",u]]),P=C;_()(h,"components",{QPage:g.Z,QCard:w.Z,QCardSection:y.Z,QForm:q.Z,QInput:v.Z,QCardActions:x.Z,QBtn:b.Z,QSeparator:k.Z})}}]);