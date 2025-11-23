# Projetos OpenSource
_Parâmetros de busca no github:_
`lang:go`
`stars:>1000`
`is:issue is:open`
`pushed:>2025-01-01`

_Sem classificação ainda:
- [HashiCorp Vault](https://github.com/hashicorp/vault) `408mb` `500k LOC`
- [HashiCorp Terraform](https://github.com/hashicorp/terraform) `396mb` `480k LOC`
- [Etcd](https://github.com/etcd-io/etcd) `116mb` `187k LOC` 
- [Prometheus](https://github.com/prometheus/prometheus) `299mb` `233k LOC` 
- [Grafana Loki](https://github.com/grafana/loki) `1GB` `9M LOC`
- [Dgraph](https://github.com/dgraph-io/dgraph)
- [MinIO](https://github.com/minio/minio)
- [Open Policy Agent](https://github.com/open-policy-agent/opa)
- [Jaeger](https://github.com/jaegertracing/jaeger)
- [Tilt](https://github.com/tilt-dev/tilt)
- [Syncthing](https://github.com/syncthing/syncthing)
- [ethereum](https://github.com/ethereum/go-ethereum)

## CLI
- [gitea](https://github.com/go-gitea/gitea) `370mb` `3.3M LOC`
- [fzf](https://github.com/junegunn/fzf) `10mb` `20k LOC`
- [Lazygit](https://github.com/jesseduffield/lazygit)  `184mb` `4.5M LOC`

## Networks
- [Caddy Server](https://github.com/caddyserver/caddy) `26MB` `60k LOC`
- [Traefik](https://github.com/traefik/traefik) `137MB` `150k LOC`
- [v2ray-core](https://github.com/v2ray/v2ray-core) `87mb` `65k LOC`
- [hysteria](https://github.com/apernet/hysteria) `26mb` `<20k LOC`
- [nsq](https://github.com/nsqio/nsq) `19mb` `<20k LOC`
- [NATS](https://github.com/nats-io/nats-server) `69mb` `300k LOC`
- Cloudflare Tools
	- [cloudflared](https://github.com/cloudflare/cloudflared) `86mb` `670k LOC`
	- [cloudflare-go](https://github.com/cloudflare/cloudflare-go) `104mb` `500k LOC`

## Containers
- [Docker(Moby)](https://github.com/moby/moby) `287mb` `1.6M LOC`
	- [dive](https://github.com/wagoodman/dive)
- [Kubernetes](https://github.com/kubernetes/kubernetes) `1.36GB` `3.6M LOC`
	- [Helm](https://github.com/helm/helm)
	- [CRI-O](https://github.com/cri-o/cri-o)
	- [Kubeshark](https://github.com/kubeshark/kubeshark)
- [Podman](https://github.com/containers/podman) `166mb` `1.4M LOC`
- [Skopeo](https://github.com/containers/skopeo) `62mb` `~500k LOC` 👀

## Web
- [Gin](https://github.com/gin-gonic/gin) `12mb` `20k LOC` 👀👀👀
- [fiber](https://github.com/gofiber/fiber) `246mb` `<80k LOC`
- [echo](https://github.com/labstack/echo) `8.5mb` `20k LOC`
- [beego](https://github.com/beego/beego) `16mb` `55k LOC`
- [iris](https://github.com/kataras/iris) `30mb` `<80k LOC`
- [Hugo](https://github.com/gohugoio/hugo)  `165mb` `160k LOC`
- [go-micro](https://github.com/micro/go-micro) `18.7mb` `44k LOC`
- [Zincsearch](https://github.com/zincsearch/zincsearch) `18mb` `140k LOC`
- [Pocketbase](https://github.com/pocketbase/pocketbase) `141mb` `92k LOC`
- [Gorilla](https://github.com/gorilla) `x mb` `n LOC` (Projeto finalizado)

## Legais/Utilitários
> Observação: estes projetos são bem simples (no sentido de não serem tão grandiosos quanto os outros, mas ainda assim úteis e importantes). Eles servem mais para familiarização com o ecosistema da linguagem.

- [gocolly](https://github.com/gocolly/colly)
- [Rclone](https://github.com/rclone/rclone)
- [bettercap](https://github.com/bettercap/bettercap)
- [Algoritmos](https://github.com/TheAlgorithms/Go)
- [transfersh](https://github.com/dutchcoders/transfer.sh)
- [ebiten](https://github.com/hajimehoshi/ebiten)
- [robotgo](https://github.com/go-vgo/robotgo)
- [tile38](https://github.com/tidwall/tile38)
- [KrillinAI](https://github.com/krillinai/KrillinAI)

# Projetos tutoriais em Go (para aprendizado prático):

[Repositório](https://github.com/practical-tutorials/project-based-learning?tab=readme-ov-file#go) com projetos tutoriais em em várias linguagens. Aqui estão os mais interessantes que eu filtrei em Go de lá:
- [Container](https://www.infoq.com/articles/build-a-container-golang/)
- [Container(Video)](https://www.youtube.com/watch?v=8fi7uSYlOdc)
- [Regex Engine](https://rhaeguard.github.io/posts/regex/)
- [Shell](https://blog.init-io.net/post/2018/07-01-go-unix-shell/)
- [TerminalEmulator](https://ishuah.com/2021/03/10/build-a-terminal-emulator-in-100-lines-of-go/)
- [LoadBalancet](https://kasvith.me/posts/lets-create-a-simple-lb-go/)
- [VideoCodec](https://github.com/kevmo314/codec-from-scratch)
- [Redis](https://www.build-redis-from-scratch.dev/en/introduction)

💡[Playlist com 59 projetos beginner friendly](https://www.youtube.com/playlist?list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9). Dá para fazer um por dia.

# 📌 TODO:
1. Organizar os projetos em categorias. (Em progresso⏳)
2. Ordenar os projetos conforme o tamanho tanto em mb quanto em linhas de código, utilizando a ferramenta cloc. (Em progresso⏳)
   (Utilizar a [API](https://api.github.com/repos/OWNER/REPO) do github) 
3. Selecionar/Peneirar os projetos mais interessantes e fazer uma análise deles.
	1. A análise consiste em ler o código linha por linha, apagar os comentários e substituir pelos nossos, explicando o que cada trecho do código faz.
	2. Isso só será feito com projetos interessantes(obviamente) e pequenos, pois não é viável no momento estudar um projeto muito grande.
