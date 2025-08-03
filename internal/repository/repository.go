package repository

// type NoteRepository struct {
// 	SiteConfig            config.SiteConfig
// 	NotesByPath           map[string]*models.ParsedNote
// 	UniqueNotesByFilename map[string]*models.ParsedNote
// 	DuplicateFilenames    map[string][]*models.ParsedNote
// 	LandingPage           *models.ParsedNote
// 	IdToNote              map[int64]*models.ParsedNote

// 	NotesPathProximityMap map[string][]*models.ScannedNote
// 	ImagePathProximityMap map[string][]string
// 	mu                    sync.RWMutex
// }

// func NewNoteRepository(siteConfig config.SiteConfig, scannedNotes []*models.ScannedNote, imagePaths []string) *NoteRepository {
// 	imagePathProximityMap := buildImagesPathProximityMap(imagePaths)
// 	notesPathProximityMap := buildNotesPathProximityMap(scannedNotes)

// 	return &NoteRepository{
// 		SiteConfig:            siteConfig,
// 		NotesByPath:           make(map[string]*models.ParsedNote),
// 		UniqueNotesByFilename: make(map[string]*models.ParsedNote),
// 		DuplicateFilenames:    make(map[string][]*models.ParsedNote),
// 		IdToNote:              make(map[int64]*models.ParsedNote),
// 		NotesPathProximityMap: notesPathProximityMap,
// 		ImagePathProximityMap: imagePathProximityMap,
// 	}
// }

// func (nr *NoteRepository) AddNote(note *models.ParsedNote) {
// 	nr.mu.Lock()
// 	defer nr.mu.Unlock()

// 	nr.populateAllRelativePaths(note)
// 	nr.IdToNote[note.ID] = note
// 	filename := note.FileName

// 	if existing, exists := nr.UniqueNotesByFilename[filename]; exists {
// 		if existing != nil {
// 			nr.DuplicateFilenames[filename] = []*models.ParsedNote{existing, note}
// 			delete(nr.UniqueNotesByFilename, filename)
// 		} else {
// 			nr.DuplicateFilenames[filename] = append(nr.DuplicateFilenames[filename], note)
// 		}
// 	} else {
// 		nr.UniqueNotesByFilename[filename] = note
// 	}
// }

// func (nr *NoteRepository) populateAllRelativePaths(note *models.ParsedNote) {
// 	relativePath := note.RelativePath
// 	if relativePath == "" {
// 		return
// 	}

// 	// Split the path into its component parts.
// 	// e.g., "Books/Comedy/Modern/BillBurr" -> ["Books", "Comedy", "Modern", "BillBurr"]
// 	components := strings.Split(relativePath, "/")

// 	// Iterate through the components to create every possible "tail" of the path.
// 	for i := 0; i < len(components)-1; i++ {
// 		// Join the remaining components to form a sub-path.
// 		// Iteration 1: "Books/Comedy/Modern/BillBurr"
// 		// Iteration 2: "Comedy/Modern/BillBurr"
// 		// Iteration 3: "Modern/BillBurr"
// 		// Iteration 4: "BillBurr"
// 		subPath := strings.Join(components[i:], "/")

// 		// Add the sub-path to the map, pointing to the original note.
// 		nr.NotesByPath[subPath] = note
// 	}
// }

// func (nr *NoteRepository) GetAllNotesSlice() []*models.ParsedNote {
// 	nr.mu.Lock()
// 	defer nr.mu.Unlock()

// 	notes := make([]*models.ParsedNote, 0, len(nr.NotesByPath))
// 	for _, note := range nr.NotesByPath {
// 		notes = append(notes, note)
// 	}
// 	return notes
// }

// func (wr *NoteRepository) AddBacklink(targetNote *models.ParsedNote, sourceNote *models.ParsedNote, pseudoName string) {
// 	wr.mu.Lock()
// 	defer wr.mu.Unlock()

// 	for _, link := range targetNote.LinkedFrom {
// 		if link.URL == sourceNote.URL {
// 			return // Already exists
// 		}
// 	}

// 	htmlFileName := utils.Slugify(sourceNote.FileName) + ".html"
// 	targetNote.LinkedFrom = append(targetNote.LinkedFrom, models.Link{
// 		Title:       sourceNote.Title,
// 		URL:         sourceNote.URL,
// 		FullPath:    sourceNote.FullPath,
// 		RawFileName: htmlFileName,
// 		PseudoName:  pseudoName,
// 	})
// }

// func (tc *NoteRepository) ResolveWikilink(wikilink, parentDir string) (*models.ParsedNote, bool) {

// 	// Resolve unique link
// 	if note, ok := tc.UniqueNotesByFilename[wikilink]; ok {
// 		return note, true
// 	}

// 	// Resolve duplicate links
// 	if duplicates, exists := tc.DuplicateFilenames[wikilink]; exists && len(duplicates) > 0 {
// 		// Is it at root?
// 		for _, note := range duplicates {
// 			if note.ParentDir == tc.SiteConfig.InputDirectory {
// 				return note, true
// 			}
// 		}

// 		// Is it in the same dir?
// 		for _, note := range duplicates {
// 			if note.ParentDir == parentDir {
// 				return note, true
// 			}
// 		}

// 		// Pick shortest path by proximity map
// 		if proximitySlice, ok := tc.NotesPathProximityMap[wikilink]; ok && len(proximitySlice) > 0 {
// 			if note, ok := tc.IdToNote[proximitySlice[0].ID]; ok {
// 				return note, true
// 			}
// 		}
// 	}

// 	// Resolve direct link
// 	if strings.Contains(wikilink, string(filepath.Separator)) {
// 		relativePath := filepath.Clean(filepath.Join(parentDir, wikilink))
// 		if note, ok := tc.NotesByPath[relativePath]; ok {
// 			return note, true
// 		}
// 	}

// 	// var noteID int64
// 	// var firstNote *models.ScannedNote
// 	// isInRootDir := false
// 	// isBareFilename := !strings.Contains(wikilink, string(filepath.Separator))
// 	// proximitySlice := tc.NotesPathProximityMap[wikilink]

// 	// if len(proximitySlice) > 0 {
// 	// 	firstNote = proximitySlice[0]
// 	// 	isInRootDir = firstNote.ParentDir == tc.SiteConfig.InputDirectory
// 	// }

// 	// if isBareFilename && isInRootDir {
// 	// 	noteID = firstNote.ID
// 	// 	if note, ok := tc.IdToNote[noteID]; ok {
// 	// 		return note, true
// 	// 	}
// 	// }

// 	// for _, note := range proximitySlice {
// 	// 	if note.ParentDir == parentDir || strings.Contains(note.RelativePath, parentDir) {
// 	// 		noteID = note.ID
// 	// 		break
// 	// 	}
// 	// }

// 	// if note, ok := tc.IdToNote[noteID]; ok {
// 	// 	return note, true
// 	// }

// 	return nil, false
// }

// func (tc *NoteRepository) GetAllByPath() map[string]*models.ParsedNote {
// 	tc.mu.RLock()
// 	defer tc.mu.RUnlock()

// 	result := make(map[string]*models.ParsedNote, len(tc.NotesByPath))
// 	maps.Copy(result, tc.NotesByPath)
// 	return result
// }

// func (tc *NoteRepository) GetAllByID() map[int64]*models.ParsedNote {
// 	tc.mu.RLock()
// 	defer tc.mu.RUnlock()

// 	result := make(map[int64]*models.ParsedNote, len(tc.IdToNote))
// 	maps.Copy(result, tc.IdToNote)
// 	return result
// }

// func (tc *NoteRepository) Count() int {
// 	tc.mu.RLock()
// 	defer tc.mu.RUnlock()
// 	return len(tc.NotesByPath)
// }

// func (tc *NoteRepository) GetByPath(path string) (*models.ParsedNote, bool) {
// 	tc.mu.RLock()
// 	defer tc.mu.RUnlock()
// 	note, ok := tc.NotesByPath[path]
// 	return note, ok
// }

// func (tc *NoteRepository) GetByFilename(filename string) (*models.ParsedNote, bool) {
// 	tc.mu.RLock()
// 	defer tc.mu.RUnlock()
// 	note, ok := tc.UniqueNotesByFilename[filename]
// 	return note, ok
// }

// func (tc *NoteRepository) GetByID(id int64) (*models.ParsedNote, bool) {
// 	tc.mu.RLock()
// 	defer tc.mu.RUnlock()
// 	note, ok := tc.IdToNote[id]
// 	return note, ok
// }

// func buildImagesPathProximityMap(imagePaths []string) map[string][]string {
// 	proximityMap := make(map[string][]string)

// 	for _, fullImagePath := range imagePaths {
// 		pathBase := filepath.Base(fullImagePath)
// 		proximityMap[pathBase] = append(proximityMap[pathBase], fullImagePath)
// 	}

// 	for _, proximitySlice := range proximityMap {
// 		slices.SortFunc(proximitySlice, func(a string, b string) int {
// 			depthA := strings.Count(a, string(filepath.Separator))
// 			depthB := strings.Count(b, string(filepath.Separator))

// 			if depthA < depthB {
// 				return -1
// 			}
// 			if depthA > depthB {
// 				return 1
// 			}
// 			return strings.Compare(a, b)
// 		})
// 	}

// 	return proximityMap
// }

// func buildNotesPathProximityMap(notes []*models.ScannedNote) map[string][]*models.ScannedNote {
// 	proximityMap := make(map[string][]*models.ScannedNote)

// 	for _, note := range notes {
// 		notePathBase := filepath.Base(note.FullPath)
// 		notePathBase = strings.TrimRight(notePathBase, filepath.Ext(notePathBase))
// 		proximityMap[notePathBase] = append(proximityMap[notePathBase], note)
// 	}

// 	for _, proximitySlice := range proximityMap {
// 		slices.SortFunc(proximitySlice, func(noteA *models.ScannedNote, noteB *models.ScannedNote) int {
// 			depthA := strings.Count(noteA.FullPath, string(filepath.Separator))
// 			depthB := strings.Count(noteB.FullPath, string(filepath.Separator))

// 			if depthA < depthB {
// 				return -1
// 			}
// 			if depthA > depthB {
// 				return 1
// 			}
// 			return strings.Compare(noteA.FullPath, noteB.FullPath)
// 		})
// 	}

// 	return proximityMap
// }
