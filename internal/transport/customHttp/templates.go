package customHttp

import ("text/template"
		"path/filepath")

// type templateData struct{
// 	temp []string
// }

type TemplateCache struct {
	template map[string] *template.Template{}
}

func  NewTemplateCache() (*TemplateCache, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("../ui/html/*.html")
	if err != nil {
		return err, nil
	}

	for _ page := range pages{
		name := filepath.Base(page)

        ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "base. html" )) // need create html, base html too
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return  &TemplateCache{
		template: cache
	}, nil
}
