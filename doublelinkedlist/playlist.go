package main

import(
	"fmt"
)

type song struct{
	name string
	artist string
	next *song
	previous *song
}

type playlist struct{
	name string
	head *song
	tail *song
	nowPlaying *song
	length int
	
}

// create the playlist

func createPlaylist(name string) *playlist{
	return &playlist{
		name: name,
	}
}

// add songs to the playlist
func (p *playlist) addSong(name, artist string) error{
	s := &song{
		name: name,
		artist: artist,
	}
	
	if p.head == nil{
		p.head = s
	}else{
		currentNode := p.tail
		currentNode.next = s
		s.previous = p.tail
	}
	// we keep track of the tail to avoid the costly big O(n)
	p.tail = s
	return nil
}

// show all songs
func (p *playlist) showAllSongs() *song{
	currentNode := p.head
	if currentNode == nil{
		fmt.Println("playlist is empty")
		return nil
	}else{
		for currentNode.next != nil{
			currentNode = currentNode.next
			fmt.Printf("%+v\n", *currentNode)
		}
	}
	return nil
}

// start playing
func (p *playlist) startPlaying() *song{
	p.nowPlaying = p.head
	return p.nowPlaying
}

// go to the next song
func (p *playlist) nextSong() *song{
	p.nowPlaying = p.nowPlaying.next
	return p.nowPlaying
}

//play previous song
func (p *playlist) previousSong() *song{
	p.nowPlaying = p.nowPlaying.previous
	return p.nowPlaying
}



// add a song at a given position
func (p *playlist) insertAtPosition(name, artist string, position int) error{
	s := &song{
		name: name,
		artist: artist,
	}
	if position < 0 {
		fmt.Println("invalid position")
	}else if position == 0{
		p.head = s
		p.length++
	}else{
		i := 0
		temp := p.head
		currentNode := temp
		for temp != nil{
			if i == position{
				currentNode.next = s
			}
		}

	}
	return nil
}

func main(){
	// create playlist
	playlistName := "myplaylist"
	myPlaylist := createPlaylist(playlistName)
	fmt.Println("created playlist")
	fmt.Println()

	// add songs
	myPlaylist.addSong("Ophelia", "The Lumineers")
	myPlaylist.addSong("Shape of you", "Ed Sheeran")
	myPlaylist.addSong("Stubborn Love", "The Lumineers")
	myPlaylist.addSong("Feels", "Calvin Harris")
	fmt.Println("showing all songs in the playlist")
	myPlaylist.showAllSongs()
	fmt.Println()
// start playing
	myPlaylist.startPlaying()
	fmt.Printf("now playing: %s by %s\n",myPlaylist.nowPlaying.name,myPlaylist.nowPlaying.artist)
	fmt.Println()

	// go to next song in the playlist
	myPlaylist.nextSong()
	fmt.Println("changing to the next song....")
	fmt.Printf("now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	myPlaylist.nextSong()
	fmt.Println("changing to the next song....")
	fmt.Printf("now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	// play previous
	myPlaylist.previousSong()
	fmt.Println("going back to the last song....")
	fmt.Printf("now playing: %s by %s\n",myPlaylist.nowPlaying.name,myPlaylist.nowPlaying.artist)
	fmt.Println()

	// insert song at a position
	myPlaylist.insertAtPosition("Stubborn Love", "Stunna",1)
	myPlaylist.showAllSongs()
}