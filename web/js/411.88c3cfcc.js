"use strict";(self["webpackChunkfrontend"]=self["webpackChunkfrontend"]||[]).push([[411],{1208:(t,n,e)=>{e.d(n,{vk:()=>r,He:()=>a,Tb:()=>i,Be:()=>l});var o=e(1768);const r={cId:0,mId:0};function a(t,n){return new Promise(((e,r)=>{const a={type:"LOGIN_REQUEST",username:t,password:n};o.api.post("/mappa/login",a).then((t=>{e(t.data)})).catch((t=>{r(t)}))}))}function i(){return new Promise(((t,n)=>{o.api.get("/hc").then((n=>t(n.data))).catch((t=>{n(t)}))}))}function l(t){try{const n=atob(t),e=JSON.parse(n);if(!e.cId)throw new Error(`Invalid context data: ${JSON.stringify(e)}`);return{cId:e.cId,mId:e.mId}}catch(n){return console.error("ParseContent",n),r}}},5411:(t,n,e)=>{e.r(n),e.d(n,{default:()=>x});var o=e(3673);const r=(0,o.Uk)(" Mappa Frontend ");function a(t,n,e,a,i,l){const u=(0,o.up)("q-btn"),c=(0,o.up)("q-toolbar-title"),s=(0,o.up)("q-toolbar"),p=(0,o.up)("q-header"),d=(0,o.up)("router-view"),f=(0,o.up)("q-page-container"),h=(0,o.up)("q-layout");return(0,o.wg)(),(0,o.j4)(h,{view:"lHh Lpr lFf"},{default:(0,o.w5)((()=>[(0,o.Wm)(p,{elevated:""},{default:(0,o.w5)((()=>[(0,o.Wm)(s,null,{default:(0,o.w5)((()=>[(0,o.Wm)(u,{flat:"",dense:"",round:"",icon:"menu","aria-label":"Menu"}),(0,o.Wm)(c,null,{default:(0,o.w5)((()=>[r])),_:1}),(0,o.Wm)("div",null,[(0,o.Wm)(u,{title:t.proxyHCTitle,flat:"",dense:"",icon:t.proxyIcon,label:t.proxyHC},null,8,["title","icon","label"])])])),_:1})])),_:1}),(0,o.Wm)(f,null,{default:(0,o.w5)((()=>[(0,o.Wm)(d)])),_:1})])),_:1})}var i=e(1208),l=function(t,n,e,o){function r(t){return t instanceof e?t:new e((function(n){n(t)}))}return new(e||(e=Promise))((function(e,a){function i(t){try{u(o.next(t))}catch(n){a(n)}}function l(t){try{u(o["throw"](t))}catch(n){a(n)}}function u(t){t.done?e(t.value):r(t.value).then(i,l)}u((o=o.apply(t,n||[])).next())}))};const u=(0,o.aZ)({name:"SingleLayout",components:{},data(){return{proxyHC:"...",proxyHCTitle:"",proxyIcon:"link"}},mounted(){return l(this,void 0,void 0,(function*(){try{let t=yield(0,i.Tb)();console.log("HC",t),this.proxyHCTitle=`Server: ${t.mappa_server.status}`,this.proxyHC="HEALTHY"==t.status?"OK":"Indisponível",this.proxyIcon="OK"==this.proxyHC?"link":"link_off"}catch(t){console.error("HC",t),this.proxyHCTitle=t.message,this.proxyHC="Em erro",this.proxyIcon="link_off"}this.proxyHC=`Proxy: ${this.proxyHC}`}))}});var c=e(3066),s=e(3812),p=e(9570),d=e(6114),f=e(3747),h=e(2652),y=e(7518),m=e.n(y);u.render=a;const x=u;m()(u,"components",{QLayout:c.Z,QHeader:s.Z,QToolbar:p.Z,QBtn:d.Z,QToolbarTitle:f.Z,QPageContainer:h.Z})}}]);