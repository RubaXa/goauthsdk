package button

import (
	"github.com/rubaxa/oauth-mailru.go"
)

type Options struct {
	Text         string
	Href         string
	Target       string
	HostStyle    string
	TextStyle    string
	PrimaryColor string
	TextColor    string
	Size         string
	Width        string
	BorderRadius string
	WithoutIcon  bool
	RedirectURI  string
}

const (
	PrimaryColor = "#005ff9"
	TextColor    = "#ffffff"
	DefaultSize  = "40px"
)

func Render(s oauth.Service, o Options) string {
	size := or(o.Size, DefaultSize)
	href := or(o.Href, s.GetAuthURL().SetRedirectURI(o.RedirectURI).String())

	return `<a
		href="` + href + `"
		target="` + or(o.Target, "_self") + `"
		style="
			width: ` + or(o.Width, "auto") + `;
			text-decoration: none;
			position: relative;
			display: inline-block;
			cursor: pointer;
			color: ` + or(o.TextColor, TextColor) + `;
			background-color: ` + or(o.PrimaryColor, PrimaryColor) + `;
			border-radius: ` + or(o.BorderRadius, "2px") + `;
			font-size: ` + size + `;
			line-height: ` + size + `;
			font-family: Arial, Helvetica, sans-serif;
			-webkit-font-smoothing: antialiased;
			-moz-osx-font-smoothing: grayscale;
			padding: 0 .4em 0 ` + choose(o.WithoutIcon, ".4", "1") + `em;
			` + or(o.HostStyle, "") + `
		">` + choose(o.WithoutIcon, "", `<span style="
				display: inline-block;
				width: .8em;
				height: .8em;
				position: absolute;
				left: 0.13em;
				top: 50%;
				margin-top: -.4em;
			">
				<svg style="width: 100%; height: 100%;" version="1.1" viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg"><g fill="none" fill-rule="evenodd"><g transform="translate(-285 -224)" fill="#FF9E00"><g transform="translate(16 170)"><g transform="translate(261 50)"><g transform="translate(8 4)"><path d="m16 7c-4.9626 0-9 4.0374-9 9.0001 0 4.9628 4.0374 9.0002 9 9.0002 1.8245 0 3.5833-0.54958 5.0695-1.5654l-1.1796-1.3714c-1.1533 0.74187-2.4975 1.1423-3.8899 1.1423-3.9732 0-7.2055-3.2324-7.2055-7.2057 0-3.9733 3.2324-7.2057 7.2055-7.2057 3.9732 0 7.2056 3.2324 7.2056 7.2057 0 0.52327-0.058737 1.0435-0.16976 1.5496-0.23099 0.94723-0.89141 1.2402-1.3988 1.2005-0.50034-0.040291-1.0947-0.39702-1.0986-1.278v-1.4721c0-2.5024-2.0359-4.5384-4.5384-4.5384-2.5023 0-4.5384 2.036-4.5384 4.5384 0 2.5026 2.0361 4.5385 4.5384 4.5385 1.2682 0 2.4159-0.52355 3.2402-1.365 0.4854 0.7752 1.2892 1.2879 2.253 1.3653 0.080353 0.006451 0.16059 0.0096199 0.24061 0.0096199 0.65307 0 1.2899-0.21328 1.8153-0.61222 0.54233-0.41168 0.93714-0.99685 1.1584-1.7048 0.035367-0.11306 0.10022-0.37506 0.10095-0.37852h-2.829e-4c0.14758-0.63899 0.19177-1.2285 0.19177-1.8529 0-4.9626-4.0374-9.0001-9-9.0001zm-2.744 9.0001c0-1.513 1.231-2.7439 2.744-2.7439 1.513 0 2.744 1.231 2.744 2.7439 0 1.5131-1.2309 2.7441-2.744 2.7441-1.513 0-2.744-1.231-2.744-2.7441z"/></g></g></g></g></g></svg>
			</span>`) +
		`<span style="font-size: .4em;` + or(o.TextStyle, "") + `">` + or(o.Text, "Mail.Ru") + `</span>
		</a>`
}

func or(v, d string) string {
	if v == "" {
		return d
	}

	return v
}

func choose(t bool, a, b string) string {
	if t {
		return a
	}

	return b
}
