package rag

import (
	"log"
	"path/filepath"
)

func LoadSources(fp string) {
	r, err := NewRag()
	if err != nil {
		log.Fatal("rag: unable to instanciate rag: " + err.Error())
	}

	err = r.createCollection()
	if err != nil {
		log.Fatal("rag: unable to create collection in db: " + err.Error())
	}
	log.Println("rag: collection created")

	pdfs, err := filepath.Glob(filepath.Join(fp, "*.pdf"))
	if err != nil {
		log.Fatal("rag: error retriving pdfs")
	}

	for _, pdf := range pdfs {
		log.Println("rag: now embedding " + pdf)
		doc, err := r.embedPDF(pdf)
		if err != nil {
			log.Fatal("rag: unable to processs pdf: " + err.Error())
		}

		for _, emb := range doc {
			err = r.insertEmbeddings(&emb)
			if err != nil {
				log.Fatal("rag: unable to insert embeded file " + emb.Id + ": " + err.Error())
			}
			//log.Println("rag: added embeded file: " + emb.Id)
		}
	}

	log.Println("rag: collection updated!")
}
