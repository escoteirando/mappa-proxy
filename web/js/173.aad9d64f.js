"use strict";(self["webpackChunkfrontend"]=self["webpackChunkfrontend"]||[]).push([[173],{1208:(t,e,n)=>{n.d(e,{vk:()=>r,He:()=>a,pT:()=>i,Tb:()=>l,Be:()=>u});var o=n(1768);const r={cId:0,mId:0};function a(t,e){return new Promise(((n,r)=>{const a={type:"LOGIN_REQUEST",username:t,password:e};o.api.post("/mappa/login",a).then((t=>{n(t.data)})).catch((t=>{r(t)}))}))}function i(t,e,n){return new Promise(((r,a)=>{const i={cId:parseInt(`${t}`),mId:parseInt(`${e}`),msg:n};console.log("Sending auth",i),o.api.post("/tg/pub",i).then((t=>{console.log("SendAuthToChat",t),r(t)})).catch((t=>{console.error("SendAuthToChat",t),a(t)}))}))}function l(){return new Promise(((t,e)=>{o.api.get("/hc").then((e=>t(e.data))).catch((t=>{e(t)}))}))}function u(t){try{const e=atob(t),n=JSON.parse(e);if(!n.cId)throw new Error(`Invalid context data: ${JSON.stringify(n)}`);return{cId:n.cId,mId:n.mId}}catch(e){return console.error("ParseContent",e),r}}},4173:(t,e,n)=>{n.r(e),n.d(e,{default:()=>H});var o=n(3673),r=n(2323);function a(t,e,n,a,i,l){const u=(0,o.up)("q-btn"),c=(0,o.up)("q-toolbar-title"),s=(0,o.up)("q-toolbar"),p=(0,o.up)("q-header"),d=(0,o.up)("router-view"),f=(0,o.up)("q-page-container"),h=(0,o.up)("q-layout");return(0,o.wg)(),(0,o.j4)(h,{view:"lHh Lpr lFf"},{default:(0,o.w5)((()=>[(0,o.Wm)(p,{elevated:""},{default:(0,o.w5)((()=>[(0,o.Wm)(s,null,{default:(0,o.w5)((()=>[(0,o.Wm)(u,{flat:"",dense:"",round:"",icon:"menu","aria-label":"Menu"}),(0,o.Wm)(c,{title:t.lastUpdate},{default:(0,o.w5)((()=>[(0,o.Uk)(" Mappa Frontend "+(0,r.zw)(t.version),1)])),_:1},8,["title"]),(0,o._)("div",null,[(0,o.Wm)(u,{title:t.proxyHCTitle,flat:"",dense:"",icon:t.proxyIcon,label:t.proxyHC},null,8,["title","icon","label"])])])),_:1})])),_:1}),(0,o.Wm)(f,null,{default:(0,o.w5)((()=>[(0,o.Wm)(d)])),_:1})])),_:1})}var i=n(1208);const l={major:0,minor:0,feature:6},u=new Date(1635079420369);var c=function(t,e,n,o){function r(t){return t instanceof n?t:new n((function(e){e(t)}))}return new(n||(n=Promise))((function(n,a){function i(t){try{u(o.next(t))}catch(e){a(e)}}function l(t){try{u(o["throw"](t))}catch(e){a(e)}}function u(t){t.done?n(t.value):r(t.value).then(i,l)}u((o=o.apply(t,e||[])).next())}))};const s=(0,o.aZ)({name:"SingleLayout",components:{},data(){return{proxyHC:"...",proxyHCTitle:"",proxyIcon:"link",version:`v${l.major}.${l.minor}.${l.feature}`,lastUpdate:`Último build: ${u.toLocaleString()}`}},mounted(){return c(this,void 0,void 0,(function*(){try{let t=yield(0,i.Tb)();console.log("HC",t),this.proxyHCTitle=`Server: ${t.mappa_server.status}`,this.proxyHC="HEALTHY"==t.status?"OK":"Indisponível",this.proxyIcon="OK"==this.proxyHC?"link":"link_off"}catch(t){console.error("HC",t),this.proxyHCTitle=t.message,this.proxyHC="Em erro",this.proxyIcon="link_off"}this.proxyHC=`Proxy: ${this.proxyHC}`}))}});var p=n(4260),d=n(3066),f=n(3812),h=n(9570),m=n(3187),y=n(3747),v=n(2652),w=n(7518),x=n.n(w);const C=(0,p.Z)(s,[["render",a]]),H=C;x()(s,"components",{QLayout:d.Z,QHeader:f.Z,QToolbar:h.Z,QBtn:m.Z,QToolbarTitle:y.Z,QPageContainer:v.Z})}}]);