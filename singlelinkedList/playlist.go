package main

import(
	"fmt"
)

type song struct{
	name string
	artist string
	next *song
}

type playlist struct{
	name string
	head *song
	nowPlaying *song
}

// create playlist

func createPlaylist(name string) *playlist{
	return &playlist{
		name: name,
	}
}

/*
The addSong() method first checks whether the linked list is empty if not it then iterates 
through the linked list by using the pointer to next song until currentSong.next becomes nil.
 When the next song is nil, it means we are at the end of the linked list. 
Now we add a new song by linking it to the current song.
*/

func (p *playlist) addSong(name, artist string) error{
	s := &song{
		name: name,
		artist: artist,
	}
	if p.head == nil{
		p.head = s
	}else{
		currentNode := p.head

		for currentNode.next != nil{
			currentNode = currentNode.next
		}
		currentNode.next = s
	}
	return nil
}

// show all songs

func (p *playlist) showAllSongs() error{
	currentNode := p.head
	if currentNode == nil{
		fmt.Println("playlist is empty")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil{
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
	return nil
}
// start playing songs from the start
func (p *playlist) startPlaying() *song{
	p.nowPlaying = p.head
	return p.nowPlaying
}
// go to the next song and play
func (p *playlist) nextSong() *song{
	p.nowPlaying = p.nowPlaying.next
	return p.nowPlaying
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
	fmt.Println("changing to the next song...")
	fmt.Printf("now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	myPlaylist.nextSong()
	fmt.Println("changing to the next song....")
	fmt.Printf("now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()
}