Read "Understanding the Linux Kernel" page cache chapter

<img width="1698" height="1362" alt="image" src="https://github.com/user-attachments/assets/9335ebb3-2a3a-46d9-8d3f-97912551fb62" />

## How is this implemented:

The Linux Page cache maintains a radix tree for every process address space. For this radix tree it maintains two lists

We have both active and inactive list to avoid Scan resistance problem, otherwise when we scan a large file, it would clear whole cache


<img width="1387" height="447" alt="image" src="https://github.com/user-attachments/assets/26c74587-4f40-4e79-92c8-95d1807e75b3" />

The active list specifies the hot list and the inactive list specifies the cold list.

Pages are initially inserted into the inactive list and when they are accessed twice, they are moved to the head of the active list. 

Ref: https://arpitbhayani.me/blogs/2q-cache/ and https://arpitbhayani.me/blogs/midpoint-insertion-caching-strategy/

### Simplified 2Q

Simplified 2Q algorithm works with two buffers: the primary LRU buffer - Am and a secondary FIFO buffer - A1. New faulted pages first go to the secondary buffer A1 and then when the page is referenced again, it moves to the primary LRU buffer Am. This ensures that the page that moves to the primary LRU buffer is hot and indeed requires to be cached.



<img width="1200" height="628" alt="image" src="https://github.com/user-attachments/assets/5d4e4fb4-556a-4023-928e-9fbea434eb12" />


## More examples with vmtouch

<img width="2874" height="798" alt="image" src="https://github.com/user-attachments/assets/90353ff8-b7f4-4259-b7dd-483ddf30fe14" />

Ref: https://biriukov.dev/docs/page-cache/4-page-cache-eviction-and-page-reclaim/

## Refault distance

When a page is evicted, the page is removed from the inactive list and its reference from the radix tree is removed too. However, in place of the removed page, a shadow meta information is stored. When a page is accessed and if there exists a shadow information then the page is directly inserted into the active page. Otherwise, the page is stored initially in the inactive list. In this way, past access information is already used in the Linux cache. However, once the page is put into the active list, this information is lost! So, the next time the same page can be evicted again when actually evicting another page could have been beneficial. This is the first problem that we try to solve.

https://github.com/torvalds/linux/blob/master/mm/workingset.c
