
package code

var (
	languageCPP = []string{
		// # Split along class definitions
		"\nclass ",
		// # Split along function definitions
		"\nvoid ",
		"\nint ",
		"\nfloat ",
		"\ndouble ",
		// # Split along control flow statements
		"\nif ",
		"\nfor ",
		"\nwhile ",
		"\nswitch ",
		"\ncase ",
		// # Split by the normal type of lines
		"\n\n",
		"\n",
		" ",
		"",
	}

	languageGo = []string{
		// Split along function definitions
		"\nfunc ",
		"\nvar ",
		"\nconst ",
		"\ntype ",
		// Split along control flow statements
		"\nif ",
		"\nfor ",
		"\nswitch ",
		"\ncase ",
		// Split by the normal type of lines
		"\n\n",
		"\n",
		" ",
		"",
	}

	languageJava = []string{
		// Split along class definitions
		"\nclass ",
		// Split along method definitions
		"\npublic ",
		"\nprotected ",
		"\nprivate ",
		"\nstatic ",
		// Split along control flow statements
		"\nif ",
		"\nfor ",
		"\nwhile ",
		"\nswitch ",
		"\ncase ",
		// Split by the normal type of lines
		"\n\n",
		"\n",
		" ",
		"",
	}

	languageKotlin = []string{
		// Split along class definitions
		"\nclass ",
		// Split along method definitions
		"\npublic ",
		"\nprotected ",
		"\nprivate ",
		"\ninternal ",
		"\ncompanion ",
		"\nfun ",
		"\nval ",
		"\nvar ",
		// Split along control flow statements
		"\nif ",
		"\nfor ",
		"\nwhile ",
		"\nwhen ",
		"\ncase ",
		"\nelse ",
		// Split by the normal type of lines
		"\n\n",
		"\n",
		" ",
		"",
	}

	languageJavaScript = []string{
		// Split along function definitions
		"\nfunction ",
		"\nconst ",
		"\nlet ",
		"\nvar ",
		"\nclass ",
		// Split along control flow statements
		"\nif ",
		"\nfor ",
		"\nwhile ",
		"\nswitch ",
		"\ncase ",
		"\ndefault ",
		// Split by the normal type of lines
		"\n\n",
		"\n",
		" ",
		"",
	}

	languageTypeScript = []string{
		"\nenum ",
		"\ninterface ",
		"\nnamespace ",
		"\ntype ",
		// Split along class definitions
		"\nclass ",
		// Split along function definitions
		"\nfunction ",
		"\nconst ",
		"\nlet ",
		"\nvar ",
		// Split along control flow statements
		"\nif ",
		"\nfor ",
		"\nwhile ",
		"\nswitch ",
		"\ncase ",
		"\ndefault ",
		// Split by the normal type of lines
		"\n\n",
		"\n",
		" ",
		"",
	}

	languagePython = []string{
		// First, try to split along class definitions
		"\nclass ",
		"\ndef ",
		"\n\tdef ",
		// Now split by the normal type of lines
		"\n\n",
		"\n",
		" ",
		"",
	}

	languageRuby = []string{
		// Split along method definitions
		"\ndef ",
		"\nclass ",
		// Split along control flow statements
		"\nif ",
		"\nunless ",
		"\nwhile ",
		"\nfor ",
		"\ndo ",
		"\nbegin ",
		"\nrescue ",
		// Split by the normal type of lines
		"\n\n",
		"\n",
		" ",
		"",
	}

	languageRust = []string{
		// Split along function definitions
		"\nfn ",
		"\nconst ",
		"\nlet ",
		// Split along control flow statements
		"\nif ",
		"\nwhile ",
		"\nfor ",
		"\nloop ",
		"\nmatch ",
		"\nconst ",
		// Split by the normal type of lines
		"\n\n",
		"\n",
		" ",
		"",
	}

	languageScala = []string{
		// Split along class definitions
		"\nclass ",
		"\nobject ",