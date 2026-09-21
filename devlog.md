# Devlog

### Sep 19, 2026
Decided to just start this project and figure things out as it progresses. 
I've been putting this off for a long time, worried about choosing the wrong approach.
Initially, I decided to use standard libraries for everything. It took a lot of time to understand and set things up.
Then I decided to go for some minimal libraries like chi to get started with programming the server. 
I could later shift to standard ones like net/http after I'm done with the mission of this project. I could still use net/http along with chi.
I won't be generating any code or text using AI. However, I'll allow myself to consult AI regarding doubts, best practices and hypotheses.

### Sep 20, 2026
Just went down the rabbit hole of knowing when to use which among panic(), log.Fatal() and os.Exit(). Also learned the trade-offs between initializing and passing down a central config object vs directly importing env variables from the runtime process.

### Sept 21, 2026
TIL its important for files to end with an empty newline. Always found it annoying when the IDE added it. Its apparently a POSIX standard practice and is a big deal for a lot of tools like grep to work. 