---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/13
- monster/size/medium
- monster/type/fiend/yugoloth
aliases: ["Ultroloth"]
NoteIcon: monster
BestiaryType: fiend (yugoloth)
SourceType: Bestiary
BookSource: Monster Manual p. 314
---
# [Ultroloth](3-Mechanics\CLI\bestiary\fiend/ultroloth.md)
*Source: Monster Manual p. 314*  

Ultroloths command the yugoloth armies of the Blood War. An ultroloth looks like a slender gray-skinned humanoid with an elongated head. Its face bears no features except for two ovoid eyes. These eyes can become sparkling pools of light that can transfix other creatures and leave them reeling and helpless.

Frequently at one another's throats, ultroloths continually scheme to enhance their own power. When not employed to fight in the Blood War, ultroloths lead yugoloth forces throughout the planes, acting as crime bosses or commanders of evil mercenary companies.

With a reputation for cruelty, ultroloths command their minions to fight while the ultroloths stay removed from combat themselves. Lesser yugoloths know their place when facing an ultroloth and respond to its summons without demand for payment.

## Yugoloths

Yugoloths are fickle fiends that inhabit the planes of Acheron, Gehenna, Hades, and Carceri. They act as mercenaries and are notorious for their shifting loyalties. They are the embodiments of avarice. Before serving under anyone's banner, a yugoloth asks the only question on its mind: *What's in it for me?*

### Spawn of Gehenna

The first yugoloths were created by a sisterhood of night hags on Gehenna. It is widely believed that Asmodeus, Lord of the Nine Hells, commissioned the work, in the hope of creating an army of fiends that were not bound to the Nine Hells. In the course of making this new army, the hags crafted four magic tomes and recorded the true names of every yugoloth they created, save one, the General of Gehenna. These tomes were called the Books of Keeping. Since knowing a fiend's true name grants power over it, the hags used the books to ensure the yugoloths' loyalty. They also used the books to capture the true names of other fiends that crossed them. It is rumored that the Books of Keeping contain the true names of a few demon lords and archdevils as well.

Petty jealousies and endless bickering caused the sisterhood to dissolve, and in the ensuing power grab, the Books of Keeping were lost or stolen. No longer indentured to anyone, the yugoloths gained independence, and they now offer their services to the highest bidder.

### Fiendish Mercenaries

Summoned yugoloths demand much for their time and loyalty. Whatever promises a yugoloth makes are quickly broken when a better opportunity presents itself. Unlike demons, yugoloths can be reasoned with, but unlike devils, they are rarely true to their word.

Yugoloths can be found anywhere, but the high cost of maintaining a yugoloth army's loyalty typically exceeds what any warlord on the Material Plane can pay. Being self-serving creatures, yugoloths quarrel among themselves constantly. A yugoloth army is more organized than a ravening horde of demons, but far less orderly and regimented than a legion of devils. Without a powerful leader to keep them in line, yugoloths fight simply to indulge their violent predilections, and only as long as it benefits them to do so.

### Back to Gehenna

When a yugoloth dies, it dissolves into a pool of ichor and reforms at full strength on the Bleak Eternity of Gehenna. Only on its native plane can a yugoloth be destroyed permanently. A yugoloth knows this and acts accordingly. When summoned to other planes, a yugoloth fights without concern for its own well-being. On Gehenna, it is more apt to retreat or plead for mercy if its demise seems imminent.

When a yugoloth is permanently destroyed, its name vanishes from every Book of Keeping. If a yugoloth is re-created by way of an unholy ritual requiring the expenditure of souls, its name reappears in the books.

### The Books of Keeping

When all four copies of the Books of Keeping disappeared, Asmodeus and the night hags lost control of their yugoloth creations. Each Book of Keeping still exists, drifting from plane to plane, where the brave and the foolish occasionally stumble upon them. A yugoloth summoned using its true name, as inscribed in the Books of Keeping, is forced to serve its summoner obediently. The yugoloth hates being controlled in this manner and isn't shy about making its displeasure known. Like a petulant child, it will follow its instructions to the letter while looking for opportunities to misinterpret them.

### The General of Gehenna

Somewhere in the brimstone wastes of Gehenna, there roams an ultroloth so strong that none contests his power: the General of Gehenna. Many yugoloths search for this great general in the hope of serving with him. They believe that service with the General of Gehenna grants power and prestige among lower planar entities.

Whatever the case, no fiend finds the General unless the General desires it. His personal name is unknown, and even the Books of Keeping contain no mention of this powerful, thoroughly evil entity.

> [!note] Variant: Yugoloth Summoning
> 
> Some yugoloths can have an action option that allows them to summon other yugoloths.
> 
> **Summon Yugoloth (1/Day).** The yugoloth chooses what to summon and attempts a magical summoning.
> 
> - An arcanaloth has a 40% chance chance of summoning one arcanaloth.  
> - A mezzoloth has a 30% chance chance of summoning one mezzoloth.  
> - A nycaloth has a 50% chance chance of summoning `dice: 1d4|avg|noform` (`1d4`) mezzoloths or one nycaloth.  
> - An ultroloth has a 50% chance chance of summoning `dice: 1d6|avg|noform` (`1d6`) mezzoloths, `dice: 1d4|avg|noform` (`1d4`) nycaloths, or one ultroloth.  
> 
> A summoned yugoloth appears in an unoccupied space within 60 feet of its summoner, does as it pleases (unless its summoner is an ultroloth, in which case it acts as an ally of its summoner), and can't summon other yugoloths. The summoned yugoloth remains for l minute, until it or its summoner dies, or until its summoner takes a bonus action to dismiss it
^variant-yugoloth-summoning

> [!quote]- A quote from Shemeshka the Marauder  
> 
> Power. We all crave it, but only a select few of us deserve it.


```statblock
"name": "Ultroloth"
"size": "Medium"
"type": "fiend"
"subtype": "yugoloth"
"alignment": "Neutral Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "153"
"hit_dice": "18d8 + 72"
"stats":
- !!int "16"
- !!int "16"
- !!int "18"
- !!int "18"
- !!int "15"
- !!int "19"
"speed": "30 ft., fly 60 ft."
"skillsaves":
  "Intimidation": !!int "9"
  "Stealth": !!int "8"
  "Perception": !!int "7"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "acid, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 17"
"languages": "Abyssal, Infernal, telepathy 120 ft."
"cr": "13"
"traits":
- "desc": "The ultroloth's innate spellcasting ability is Charisma (spell save DC\
    \ 17). The ultroloth can innately cast the following spells, requiring no material\
    \ components:\n\nAt will: [alter self](/3-Mechanics/CLI/spells/alter-self.md),\
    \ [clairvoyance](/3-Mechanics/CLI/spells/clairvoyance.md), [darkness](/3-Mechanics/CLI/spells/darkness.md),\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md)\
    \ (self only), [suggestion](/3-Mechanics/CLI/spells/suggestion.md)\n\n1/day\
    \ each: [fire storm](/3-Mechanics/CLI/spells/fire-storm.md), [mass suggestion](/3-Mechanics/CLI/spells/mass-suggestion.md)\n\
    \n3/day each: [dimension door](/3-Mechanics/CLI/spells/dimension-door.md),\
    \ [fear](/3-Mechanics/CLI/spells/fear.md), [wall of fire](/3-Mechanics/CLI/spells/wall-of-fire.md)"
  "name": "Innate Spellcasting"
- "desc": "The ultroloth has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The ultroloth's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "The ultroloth can use its Hypnotic Gaze and makes three melee attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) slashing damage, or dice:1d10 + 3|text(8)\
    \ (1d10 + 3) slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "The ultroloth's eyes sparkle with opalescent light as it targets one creature\
    \ it can see within 30 feet of it. If the target can see the ultroloth, the target\
    \ must succeed on a DC 17 Wisdom saving throw against this magic or be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ until the end of the ultroloth's next turn. The [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ target is [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned). If the target's\
    \ saving throw is successful, the target is immune to the ultroloth's gaze for\
    \ the next 24 hours."
  "name": "Hypnotic Gaze"
- "desc": "The ultroloth magically teleports, along with any equipment it is wearing\
    \ or carrying, up to 60 feet to an unoccupied space it can see."
  "name": "Teleport"
"source":
- "MM"
- "WDMM"
- "BGDIA"
- "EGW"
- "JttRC"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Ultroloth.webp"
```
^statblock

```encounter-table
name: Ultroloth
creatures:
 - 1: Ultroloth
```