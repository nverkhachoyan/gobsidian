---
questObtained: 2024-11-12
questStatus: Complete
questGiver: 
questLocationObtained: 
questSessionObtained: "[[1-Session Journals/2024-11-12.md|2024-11-12]]"
questNotes: 
questLootAvail:
  - "[[bag-of-holding|Bag of Holding]]"
questLootEarned: 250gp
NoteIcon: quest
obsidianUIMode: preview
tags:
  - quest
---

# `=this.file.name`

> [!infobox]+
> # `=this.file.name`
> ## Quest Details
> Type |  Stat |
> ---|---|
> Date Obtained | `INPUT[datePicker:questObtained]` |
> Status | `INPUT[inlineSelect(option(Not Started), option(In Progress), option(Complete)):questStatus]` |
> Quest Giver | `INPUT[suggester(optionQuery(#npc)):questGiver]` |
> Quest Location | `INPUT[suggester(optionQuery(#Category/Settlement)):questLocationObtained]` |
> Session Obtained | `INPUT[suggester(optionQuery(#journal)):questSessionObtained]` |
> Available Loot | `=this.questLootAvail` |
> Acquired Loot | `=this.questLootEarned` |

The PCs go through a entry point scene in the Outer Planes with the goddess, [[faerunian-tymora|Tymora]], the goddess of luck and fortune.

As it turns out, Tymora is a clumsy goddess and accidentally killed the PCs in their home worlds instead of summoning a different set of heroes.

### Quest Objective

- [?] Negotiate [[forgotten-realms-tymora|Tymora]] for bonuses.
	- [ ] A Feat of any PCs choice.
	- [ ] An ability score improvement right of the bat.
	- [ ] [[bag-of-holding|Bag of Holding]], like stats to the PCs built in inventory mechanic.
	- [x] Starting currency between 50 and 250gp, depending on how the PCs negotiate.

### Rewards

- [[bag-of-holding|Bag of Holding]]]
- 50-250gp.

### Scenario

- [>] The PCs initially see, hear and feel nothingness.
- [>] Sparkles begin to glisten around them and eventually expands to a white empty space with nothing in side, except the Goddess herself.

> [!readaloud]
> Before you is a being who appears to be a woman with brown hair and green eyes.  Her hair appears to sparkle and glisten as she walks around.  It is in a messy bob-like hairstyle.  She has long pointy ears and a pair of horns which curl around in a small spiral.  She is wearing a navy blue dress with golden details throughout. She has jewelry all around.
> 
> As your vision begins to process the situation you find youself in, you see her pacing back and forth.
> She says, "Oh no, what am I going to do?? Talos is going to be so mad at me!!!"
> 
> She continues on, not yet noticing your presence.

- [?] Upon noticing the PCs, she exclaims, "OH, Your here already!?!"
- [>] She introduces herself as [[faerunian-tymora|Tymora]], the goddess of fortune, also known as lady luck.
- [>] She explains that she was trying to summon heroes from another realm, but after getting distracted she accidentally caused enough misfortune on the PCs, that they were killed in their original realms and brought here instead.
- [?] She is deeply remorseful for ending their lives so quickly and will do anything to get on the good side of the PCs.  
- [?] She is unable to summon more heroes from a different realm for another 100 years in the Forgotten Realms.
