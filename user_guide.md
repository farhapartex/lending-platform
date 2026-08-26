# A Guide to the Lending Platform

Welcome. This guide explains what this platform is, what you can do on every page, and — just as
importantly — what you cannot do yet.

It assumes you have used MetaMask before and know how to approve and sign a transaction when your wallet
asks. It does **not** assume you know anything about lending, borrowing, or how a platform like this works
underneath. Every term is explained the first time it appears.

Read the first three sections before you touch anything. They will save you from the two mistakes that cost
people real money.

---

## Table of contents

1. [What this platform does](#1-what-this-platform-does)
2. [The words you need](#2-the-words-you-need)
3. [The two tokens, and why your ETH is not one of them](#3-the-two-tokens-and-why-your-eth-is-not-one-of-them)
4. [What works today](#4-what-works-today)
5. [Getting set up](#5-getting-set-up)
6. [Page by page](#6-page-by-page)
   - [Home](#home--)
   - [Welcome](#welcome-welcome)
   - [Markets](#markets-markets)
   - [Lend](#lend-lend)
   - [Borrow](#borrow-borrow)
   - [Dashboard](#dashboard-dashboard)
   - [History](#history-history)
   - [Liquidations](#liquidations-liquidations)
   - [Practice](#practice-practice)
   - [Learn](#learn-learn)
7. [Three walkthroughs](#7-three-walkthroughs)
8. [What you cannot do](#8-what-you-cannot-do)
9. [Rules that keep you out of trouble](#9-rules-that-keep-you-out-of-trouble)
10. [Glossary](#10-glossary)

---

## 1. What this platform does

Two things, and they are two sides of the same coin.

**You can lend.** You put a stable digital dollar called USDC into a shared pool. Other people borrow from
that pool and pay interest. That interest comes to you. Your money is not locked up — you can take it out
whenever there is enough left in the pool.

**You can borrow.** You lock up something you own — a token called WETH, which tracks the price of Ether —
and take out a USDC loan against it. You keep ownership of the WETH. You never sold it. When you repay the
loan plus interest, you get your WETH back.

The point of borrowing this way is that you get spendable money **without selling** what you own. If you
believe your WETH will be worth more later, selling it feels wrong. Borrowing against it does not.

There is a catch, and it is the whole reason the rest of this guide exists: **if the price of your locked-up
WETH falls far enough, a stranger is allowed to repay part of your loan and take some of your WETH.** That
is called liquidation. It is not a punishment or a bug — it is the mechanism that protects the lenders'
money. Section 9 explains how to stay well clear of it.

Nobody is in charge of the pool. There is no company holding your money, no manager approving your loan, and
no support desk that can reverse a transaction. Rules are written into software on a blockchain and run the
same way for everyone, every time. That is a genuine strength and a genuine risk, and you should understand
both before you commit anything.

---

## 2. The words you need

If you have used MetaMask you already know some of this. It is here in one place so nothing later in the
guide is a mystery.

### Blockchain

A shared record book that thousands of computers keep a copy of. When you do something, the record is added
to every copy. Nobody can quietly edit an old page. This is why actions here are **final** — there is no
"undo", and no one to ask for one.

### Wallet

An app, usually a browser extension like MetaMask, that holds the secret key proving certain money is yours.
It is not an account with us. **We never see your key, and we cannot recover it.** Lose it and the money is
gone permanently. There is no password reset.

### Address

Your public identity on the blockchain — a long string starting with `0x`. Safe to share; it is how the
system knows which balances are yours. You will see yours shortened, like `0xf39F…2266`.

### Transaction

Any action that changes the record book: sending money, depositing, borrowing. Your wallet asks you to sign
it, then it takes a few seconds to a minute to become final. Until it is final it is "pending".

### Gas

The fee the blockchain charges to process your transaction. It is paid in the network's own currency (ETH),
**not** in USDC or WETH. Two consequences worth remembering:

- You must always keep a little ETH in your wallet, or you cannot do anything at all — not even withdraw.
- **We charge no fees. Gas is not ours.** It goes to the network. We never take a cut of a deposit,
  withdrawal, borrow, or repayment.

### Token

A unit of value living on the blockchain that isn't the blockchain's own currency. Think of the blockchain as
a country and its currency (ETH) as the national money — tokens are everything else that circulates there:
company shares, digital dollars, loyalty points.

### ERC-20

A shared rulebook that most tokens follow. "ERC-20 token" just means "a token built to the standard
pattern".

Why it matters to you: because every ERC-20 token behaves the same way, your wallet knows how to show it,
and our platform knows how to accept it, without anyone writing special code for each one. It is the reason
a new token appears in MetaMask correctly without any effort.

One ERC-20 rule shapes how you will use this site. A token contract will not let anyone move your tokens
unless you have first given that specific party written permission for a specific amount. So depositing
takes **two** transactions, not one:

1. **Approve** — you tell the token: "this platform may take up to 500 USDC of mine."
2. **Deposit** — the platform takes it.

The first time you deposit or repay a given token, expect two wallet pop-ups. That is normal, not a
double-charge. Approving costs gas but moves no money. And approving is not depositing — if you approve and
then walk away, nothing has happened.

### Decimals

Blockchains cannot handle fractions, so tokens store whole numbers and remember where the decimal point
belongs. USDC uses 6 decimals, WETH uses 18. So "1 USDC" is stored as `1000000`.

You will never need to do this maths — every screen shows you human numbers. It is worth knowing only
because the two tokens use different decimals, so "1.5" of one is not comparable to "1.5" of the other.

### Price feed (oracle)

The platform needs to know what WETH is worth to judge whether your loan is safe. It cannot look this up
itself, so it reads a **price feed**: a service that publishes the current price onto the blockchain.

If that feed stops updating, the platform notices the price has gone **stale** and refuses to act on it. You
will see a warning, and actions will be disabled. That is deliberate — acting on a price that might be
hours old is how people lose money unfairly. Your balances are untouched; you simply wait.

### Interest, APY and APR

**APY** is what you earn as a lender, per year, if things stay as they are. **APR** is what you pay as a
borrower. Both are shown as percentages, and both **move constantly** — they are not promises.

They move because they are set by supply and demand in the pool. When lots of people are borrowing, rates go
up, which attracts more lenders and discourages borrowing. When few are borrowing, rates fall. Nobody sets
them by hand.

Interest accrues continuously, not monthly. Your lender balance ticks up second by second, and a borrower's
debt ticks up the same way.

### Utilization

The share of the pool currently lent out. If lenders put in 100 and borrowers took 65, utilization is 65%.
It drives the rates. There is a point — here, 80% — where rates start climbing much faster, deliberately, to
pull the pool back from being fully drained. If utilization ever hit 100%, lenders could not withdraw.

### Collateral

What you lock up to back a loan. Here, that is WETH. It stays yours and comes back when you repay. But while
it is locked, it is at risk if its price falls.

### Loan-to-value (LTV)

Your debt as a percentage of your collateral's value. Borrow 1,000 USDC against 4,000 USDC of WETH and your
LTV is 25%. Lower is safer. Three numbers matter here:

| Number | Value | Meaning |
| --- | --- | --- |
| Maximum borrow | **75%** | The most you are allowed to borrow against your collateral |
| Liquidation threshold | **80%** | Cross this and you can be liquidated |
| Recommended | **55%** | What we suggest you actually stop at |

The gap between 75% and 80% is a deliberate cushion. If borrowing were allowed right up to the liquidation
point, a tiny price wobble would liquidate you seconds after you borrowed.

### Health factor

The single number that tells you how safe your loan is. It is the same information as LTV, expressed so that
**bigger is safer** and there is one line that matters:

- **Above 1.00** — you are fine
- **Exactly 1.00** — you are at the edge
- **Below 1.00** — anyone may liquidate you

The platform groups it into four bands:

| Band | Health factor | What it means |
| --- | --- | --- |
| **Safe** | 1.50 and above | Comfortable. The price has room to move against you |
| **Caution** | 1.15 to 1.50 | Still safe, but thinner than we would recommend |
| **At risk** | 1.00 to 1.15 | A small price fall will liquidate you. Act now |
| **Liquidatable** | Below 1.00 | Anyone can liquidate you right now |

If you have no loan, there is nothing to be unsafe about, and the platform shows no health factor at all.

### Liquidation

When your health factor drops below 1.00, anyone may step in, repay part of your debt for you, and take some
of your collateral in return — plus a **5% bonus** on top, as their reward for doing it.

You lose the collateral that was taken, and the bonus. You keep the rest, and the debt they repaid is really
gone.

This is not the platform punishing you. It is the only way lenders can be sure of getting their money back,
and it is why lending here is safe enough for anyone to do. The 5% bonus exists to make sure somebody
actually bothers, quickly, before the position gets worse.

### Testnet

A practice copy of a blockchain. Same software, same behaviour, but the money is worthless play money that
you get for free.

**This matters a great deal here: this platform currently runs on a test network only.** Details in the next
section.

---

## 3. The two tokens, and why your ETH is not one of them

The platform uses **exactly two tokens**. Not dozens. Two.

| Token | Full name | What it is | Its job here | Decimals |
| --- | --- | --- | --- | --- |
| **WETH** | Wrapped Ether | A token that tracks the price of ETH, one for one | **Collateral only.** Lock it up to borrow against | 18 |
| **USDC** | USD Coin | A token designed to always be worth about one US dollar | **The lending token.** Lend it to earn, or borrow it | 6 |

Each has one job and cannot do the other's. **You cannot lend WETH, and you cannot post USDC as
collateral.** Later versions may add more markets; today there is one pair.

### Why USDC for lending

Because it holds its value. If you lend something whose price swings wildly, a good interest rate means
nothing — the price movement dwarfs it. Lending a dollar-stable token means your earnings are actually
earnings.

### Why WETH and not ETH

Here is the part that catches people out.

**ETH is not a token.** It is the blockchain's own built-in currency, and it predates the ERC-20 rulebook,
so it does not follow those rules. Our contracts are written to handle ERC-20 tokens.

**WETH is the ERC-20 version of ETH.** You lock 1 ETH into a wrapping contract and receive 1 WETH, which
behaves like a normal token. Unwrap it later and you get your ETH back. The price is the same; WETH is
simply ETH in a standard-shaped box.

### So what can you do with ETH in your wallet?

Read this carefully, because it is probably why you are here.

**ETH pays your gas — and that is essential.** Keep some. Without ETH you cannot transact at all.

**ETH cannot be deposited, lent, or used as collateral.** There is no screen on this platform that accepts
ETH. If you have only ETH, you cannot yet lend or borrow.

**To take part you need WETH, USDC, or both:**

| You want to | You need | Why |
| --- | --- | --- |
| Lend and earn interest | **USDC** | It is the only lending token |
| Borrow | **WETH** as collateral | It is the only collateral |
| Repay a loan | **USDC** | You borrowed USDC, you repay USDC |
| Do anything at all | A little **ETH** | For gas |

**On a test network, you get both tokens free from the [Practice](#practice-practice) page.** That is the
intended route, and it means you do not need to convert anything. Turning real ETH into WETH is done on
other services (a wrapping site or an exchange), not here — we have no swap or wrap feature, and none is
planned for this phase.

### The honest bit about these tokens

On the test network this platform runs on, the WETH and USDC are **not the real ones**. They are stand-ins
created for practice. They:

- cannot be sold or exchanged for anything
- cannot be moved to the real network
- may vanish without warning if the test network is reset

So any interest you earn here is not real income, and any loss is not a real loss. Treat everything you do
as a rehearsal. That is exactly what it is for.

---

## 4. What works today

This platform is being built in stages, and this is an early one. Being straight with you about what is
finished matters more than looking finished.

### You can do this now

- Browse every page and see how the whole thing fits together
- Connect your wallet and confirm you are on the right network
- Read all the learning material
- Use the price-drop simulator on the Borrow page to see how a falling price would affect a loan
- See genuinely live data in three places: **recent activity** on the Dashboard, and **past liquidations**
  and their **receipts** on the Liquidations page

### You cannot do this yet

**No action that moves money is connected.** You can open the deposit, borrow, repay and liquidate screens,
fill in an amount, and see the review panel with its warnings. But the final button will not produce a
wallet pop-up, because those screens are not yet wired to the blockchain.

Concretely, you cannot yet: lend USDC, withdraw it, deposit or withdraw collateral, borrow, repay, liquidate
anyone, or claim test tokens from the faucet.

### Numbers you should not trust yet

Most figures on screen are **sample values**, put there so the pages could be designed and reviewed. Rates,
pool totals, your position sizes, and the list of positions eligible for liquidation are all placeholders.

The three live panels are the exception. And because the service that reads blockchain history has not been
switched on yet, those panels will honestly tell you so — you will see **"Activity history is not ready
yet"** rather than a misleading empty list. That message is a feature: it distinguishes "nothing has
happened" from "we cannot see what happened yet".

**In short: everything in this guide describes how the platform is designed to work, and reading it now is
the best possible preparation. Just do not expect a deposit to go through today.**

---

## 5. Getting set up

### What you need

1. **A wallet.** MetaMask, Coinbase Wallet, or anything that connects via WalletConnect.
2. **The right network.** The platform runs on one network at a time. If your wallet is pointed elsewhere,
   a banner appears across the top and everything else hides until you switch. Your wallet will offer to
   switch for you.
3. **A little ETH** on that network, for gas.
4. **Test tokens** from the Practice page, once that is live.

### Connecting

Click **Connect wallet**, top right. Pick your wallet, approve the connection request.

That is the entire sign-up. **There is no account, no email, no password, and no profile.** Your wallet
address *is* your identity. Disconnect and reconnect any time; nothing is lost, because nothing was stored
with us in the first place.

Connecting is read-only. It lets the site see your address and balances. It does **not** let us move
anything — every movement needs a separate signature from you.

---

## 6. Page by page

Each page below is marked with what is live and what is placeholder.

---

### Home — `/`

**Status: presentation page. Everything works, the numbers are samples.**

The public front page, visible without a wallet. It explains the platform to someone who has never seen it
and gives you a way in.

You will find:

- A plain-English summary of the two things you can do
- **Live-looking protocol totals** — how much is deposited, how much borrowed, how busy the pool is. Sample
  figures for now
- **How it works**, in three steps
- **A fee summary.** Worth reading, because the answer is short: no deposit fee, no withdrawal fee, no
  borrowing fee. The only ongoing charge is that 10% of the interest borrowers pay is kept by the protocol
  rather than passed to lenders — that is how the reserve that protects lenders is built. And the 5%
  liquidation bonus, which only ever applies if your position becomes unsafe
- **Security notes**, on how the contracts were tested
- An invitation to try Practice mode first

**What to do here:** read the fee section, then go to Welcome.

---

### Welcome — `/welcome`

**Status: fully working. It is all explanatory content.**

A guided introduction for newcomers, and the best first stop. It walks through, in order:

- The two ways to put crypto to work
- Why there is no account to create
- Earning on crypto you already hold
- Borrowing without selling what you own
- The one number that tells you how you are doing — the health factor
- How to stay on the safe side
- What is worth knowing while the platform is early
- Why to try play money first

There is nothing to click that costs anything. If any term in this guide felt thin, this page covers the
same ground with pictures and examples.

---

### Markets — `/markets`

**Status: page works, figures are samples. Live rates arrive with the next stage.**

The state of the pool. This is where you check conditions **before** deciding to lend or borrow. No wallet
needed — it is public information.

| What you see | What it tells you |
| --- | --- |
| **Supply APY** | What lenders are earning right now |
| **Borrow APR** | What borrowers are paying right now |
| **Max borrow — 75%** | The most you may borrow against your collateral |
| **Liquidation at — 80%** | Where liquidation becomes possible |
| **Available liquidity** | USDC sitting in the pool, not lent out. **This is the cap on withdrawals right now** |
| **Utilization bar** | How much of the pool is lent out, and where the 80% "kink" sits |
| **WETH price ticker** | The price the platform is using, and how long ago it updated |
| **Fee disclosure** | The same short fee list as the home page |

Two things to actually use this page for:

**Before lending, look at available liquidity.** If it is low, the pool is heavily borrowed. You will earn
more, but withdrawing may have to wait until borrowers repay.

**Before borrowing, look at the borrow APR and utilization.** If utilization is near or above 80%, the rate
climbs steeply, and your loan may get more expensive while you hold it.

If the price feed has gone stale, a warning appears here first.

---

### Lend — `/lend`

**Status: screens work, amounts validate, position figures are samples. Deposits and withdrawals are not
connected yet.**

Where you put USDC in to earn interest. Three parts.

**Your position** — how much you have deposited, what it is worth now, and what you have earned. There is a
live-ticking interest counter, which is not decoration: interest really does accrue continuously.

**Deposit and withdraw** — a two-tab panel.

*Depositing* asks for an amount, with a **Max** button. It checks your wallet balance and the pool's
minimum, and tells you plainly if something is wrong instead of failing at the wallet stage. Because USDC is
an ERC-20 token, your first deposit needs the **approve-then-deposit** pair described in section 2 — the
panel shows this as two clear steps.

*Withdrawing* is the same in reverse, with one thing to understand: **you can only withdraw what is actually
in the pool.** If borrowers have taken most of it, your withdrawal is capped by available liquidity, and the
panel says so. Your money is not lost or locked away — it comes back as borrowers repay. It is the normal
trade-off for earning interest, but it is a real limit and worth knowing before you need the money in a
hurry.

**Market conditions** — the rate you are earning and why, so you do not have to leave the page.

**Risks of lending, honestly:** your USDC is lent to strangers. It is protected by their collateral and by
liquidation, which is why this is the lower-risk side. But "lower risk" is not "no risk". Withdrawal depends
on liquidity, and a bug in the software would be everyone's problem.

---

### Borrow — `/borrow`

**Status: screens and the simulator work, position figures are samples. Borrowing and repaying are not
connected yet.**

The most involved page, because borrowing is the part that can go wrong. Take it in order.

**Position health, at the top.** A gauge showing your health factor and which of the four bands you are in,
plus a bar marking the 75% borrow limit and the 80% liquidation line. If you have no loan, there is nothing
to show.

**Your collateral** — deposit or withdraw WETH.

Depositing WETH is what makes borrowing possible. Same approve-then-deposit pair.

Withdrawing collateral is where people get hurt. Taking collateral out while you owe money pushes your
health factor **down**. The page will not let you withdraw an amount that would immediately make you
liquidatable, and it shows the maximum that keeps you safe. **Do not fight that limit.** It is protecting
you from being liquidated by your own withdrawal.

**Your loan** — borrow or repay.

*Borrowing* shows a slider from zero to your maximum, with the recommended **55%** marked. As you move it,
you see the health factor you will end up with, before you commit. Push toward 75% and warnings escalate.

The slider stops at 75% of your collateral value, and the page explains where that ceiling came from. If the
pool lacks the USDC, it tells you rather than letting you submit a doomed transaction.

*Repaying* offers a partial amount or **Repay everything**. Use "everything" when closing out: debt grows
every second, so an amount typed manually is stale by the time it confirms, and you can be left owing a few
cents that keeps the loan open.

**Price drop simulator.** This one genuinely works today — it is a calculator, not a transaction. Drag a
slider to drop the WETH price by some percentage and see what happens to your health factor, and at which
price you would be liquidated.

**Use it before you borrow, every time.** It converts "I am at 60% LTV" into "a 20% price fall liquidates
me", which is the form your brain can actually act on. Crypto prices fall 20% in a day more often than
people expect.

---

### Dashboard — `/dashboard`

**Status: position figures are samples. Recent activity is live.** Needs a connected wallet.

Everything about your own money in one place. Lending and borrowing side by side.

**Portfolio overview** — total deposited, total collateral, total borrowed.

**Health and risk** — the same gauge as the Borrow page, plus a legend explaining the bands, so this page
alone tells you whether you need to act.

**Warnings**, if they apply: a risk warning when your health factor is low, and a banner if you have been
liquidated, which opens a receipt showing exactly what happened.

**Quick actions** — shortcuts to lend, borrow, repay, or add collateral.

**Your positions** — a lender card and a borrower card.

**Recent activity — this is live.** Your last few actions, read from our history service. It is honest about
its own state: if history is not being indexed yet it says so, rather than implying you have done nothing.
A link takes you to the full History page.

**Price drop drawer** — the simulator again, reachable from here.

**One thing to note:** your position figures come from the blockchain, but activity comes from our history
service. If our service has a problem, the activity panel will say so **and the rest of the page keeps
working**. That is deliberate: a problem on our side must never stop you seeing or managing your own
position.

---

### History — `/history`

**Status: the table is sample data. The detail panel is connected and will show live records once history
is being indexed.** Needs a connected wallet.

Your complete record: every deposit, withdrawal, borrow, repayment, collateral move and liquidation.

- **Filter by type** — show only borrows, only repayments, and so on
- **Filter by date** — last 7, 30 or 90 days, or everything
- **Page through** older entries
- **Click any row** for a detail panel with the full record: exact amount, your health factor immediately
  afterwards, block number, and a link to view it on a public blockchain explorer

Each entry links out to an explorer so you can confirm independently that what we show matches the
blockchain. You never have to take our word for it.

**Why the delay?** After you sign something, it takes roughly 10 to 30 seconds to appear here. The record
must be confirmed on the blockchain, then read by our service. Your transaction is not lost — check your
wallet or the explorer if you are unsure.

---

### Liquidations — `/liquidations`

**Status: the eligible list is sample data and liquidating is not connected. Past liquidations and receipts
are live.**

Public — no wallet needed to look. Two halves.

**Positions eligible now.** Loans that have fallen below a health factor of 1.00 and can be resolved by
anyone. Each row shows the borrower's health, how much debt you would repay, how much collateral you would
receive, and your bonus. You can sort by health, size, or reward.

Anyone can do this. It is not reserved for professionals, though in practice automated bots move fast.

**A safety point that matters if you ever do this:** eligibility is re-checked on the blockchain at the
moment you act. Between the list loading and your transaction landing, someone else may have got there
first, or the price may have recovered — and your transaction will fail. That is the system working
correctly, not a fault. Never treat the list as a promise.

**Liquidations already settled — this is live.** Every liquidation that has completed, newest first, showing
what was repaid, what was seized, and the bonus earned. Click **Receipt** on any row for the full record:
the health factor before, the exact price that triggered it, both parties, and the block.

Rows are flagged **"Lost money"** where the liquidation ran at a shortfall — the collateral was worth less
than the debt plus bonus, so the liquidator paid out more than they received. Rare, but real, and worth
studying before you attempt one.

**This half is the most useful page on the site for learning.** Real settled liquidations, with real
numbers, show you what actually happens when a position goes bad — far more instructive than any
explanation.

---

### Practice — `/practice`

**Status: page and guidance work. The faucet is not connected yet.**

How to try everything with free play money. Recommended before you use anything you care about.

- **Network switch card** — how to point your wallet at the test network
- **Test token faucet** — free tokens, roughly **5,000 USDC** and **2 WETH**, limited to **one request per
  day** so nobody drains it
- **What is real and what is not** — an honest split, reproduced below
- **Things to try** — suggested exercises

**Real, even in practice mode:** the interface is identical to the one you would use with real money;
rates, safety scores and liquidation rules behave exactly the same; your actions are genuine blockchain
transactions, so you experience real confirmations and gas.

**Not real:** the tokens are worthless and cannot be sold, swapped, or moved to the live network; any
interest or loss is imaginary; test networks are occasionally reset, which clears balances without notice.

**The exercise worth doing:** borrow close to the limit, then use the practice price controls to drop the
WETH price until your position becomes liquidatable, and watch a liquidation happen to you. Ten minutes here
teaches you more than this entire guide.

---

### Learn — `/learn`

**Status: fully working. It is all reference material.**

Reference documentation, with a sidebar and a table of contents on each page:

| Page | What it covers |
| --- | --- |
| **How it works** | The full mechanics of the pool, from deposit to repayment |
| **Health score** | The health factor in depth, with the tier table and worked examples |
| **Liquidation** | Exactly how liquidation runs, with a worked bonus calculation |
| **Fees** | Every charge, in full |
| **Glossary** | Every term used anywhere on the platform |
| **FAQ** | Common questions |

Terms elsewhere on the site link back here, so you can always find out what something means without losing
your place.

---

## 7. Three walkthroughs

These describe the intended flow. The reading steps work today; the signing steps wait on the next stage.

### Lending USDC to earn interest

1. Open **Markets**. Check the supply APY and available liquidity.
2. Open **Lend**. Enter an amount, or press **Max**.
3. Approve the token spend when your wallet asks. *(Costs gas, moves nothing.)*
4. Confirm the deposit.
5. Wait for confirmation, then see your position on **Lend** or **Dashboard**.
6. Withdraw whenever you like, subject to available liquidity.

**Watch out for:** stopping after the approval and thinking you have deposited. Two pop-ups, both needed.

### Borrowing USDC against WETH

1. Open **Markets**. Check the borrow APR.
2. Open **Borrow**, go to **Your collateral**, deposit WETH. Approve, then deposit.
3. Move to **Your loan**. Use the slider — **stop at the 55% recommendation**, not the 75% maximum.
4. Check the health factor the panel predicts. Aim for **1.50 or better**.
5. Open the **price drop simulator** and check what a 20% and a 30% fall would do. If either liquidates you,
   borrow less.
6. Confirm. The USDC arrives in your wallet.
7. **Check your health factor regularly from then on.** The price moves while you sleep.
8. To close out: repay with **Repay everything**, then withdraw your WETH.

**Watch out for:** borrowing the maximum. At 75% you are five percentage points from liquidation, which a
normal day's price movement can cover.

### Liquidating someone else's position

1. Open **Liquidations**. Study the **settled** list first to see what real payouts look like.
2. Look at the eligible list. Sort by reward.
3. Check the row: debt to repay, collateral to receive, bonus. Make sure you hold enough USDC.
4. Open the liquidation panel. It re-checks eligibility on the blockchain.
5. Approve the USDC spend, then confirm.
6. If it fails with "position is healthy", someone beat you to it or the price recovered. Normal.

**Watch out for:** a shortfall. Check that the collateral you receive is genuinely worth more than the USDC
you pay. The settled list flags past liquidations where it was not.

---

## 8. What you cannot do

Clear limits, so you do not go looking for features that are not there.

### Not possible by design

| You cannot | Why |
| --- | --- |
| Deposit or lend **ETH** | Only ERC-20 tokens are accepted. Use WETH |
| Lend **WETH** | WETH is collateral only |
| Post **USDC** as collateral | USDC is the lending token only |
| Use any token other than these two | One market exists today |
| Swap or wrap tokens here | We are not an exchange. Wrap ETH elsewhere |
| Borrow with no collateral | Every loan is backed. There are no credit checks and no unsecured lending |
| Borrow more than 75% of your collateral | The hard ceiling |
| Withdraw more USDC than the pool holds | Capped by available liquidity |
| Withdraw collateral that would make you liquidatable | Blocked to protect you |
| Reverse a transaction | Blockchains are final. Nobody can undo it |
| Recover a lost wallet key | We never had it |
| Get help from a support desk | There isn't one. That is what this guide is for |
| Stop someone liquidating you | If you go below 1.00, anyone may act. Repay or add collateral first |

### Not possible *yet*

| You cannot yet | Coming |
| --- | --- |
| Lend or withdraw USDC | Once the action screens are connected |
| Deposit or withdraw collateral | Same |
| Borrow or repay | Same |
| Liquidate a position | Same |
| Claim test tokens from the faucet | Same |
| See real rates, totals or position sizes | Once live figures replace the samples |
| See your real transaction history | Once the history service is switched on |
| Use real money | Test networks only, for now |

---

## 9. Rules that keep you out of trouble

The short list. If you remember nothing else:

**1. Never borrow the maximum.** 75% is allowed; 55% is sensible. That gap is your protection against a
normal bad day.

**2. Keep your health factor above 1.50.** Below 1.15 you are one modest price move from losing collateral.

**3. Run the simulator before borrowing.** "A 25% fall liquidates me" is actionable. "My LTV is 68%" is not.

**4. Always keep some ETH for gas.** Run out and you cannot repay, cannot add collateral, cannot withdraw —
precisely when you most need to.

**5. Approving is not depositing.** Two pop-ups. Finish both.

**6. Use "Repay everything" to close a loan.** Interest accrues every second; a typed amount leaves dust
behind and keeps the loan open.

**7. Check on a borrow position regularly.** Prices move while you sleep, and liquidation needs no warning.

**8. Practise first.** Free tokens, identical mechanics. Get liquidated on purpose once, on play money.

**9. Never trust a screen over the blockchain.** Every record links to a public explorer. Verify anything
that matters.

**10. Only commit what you can afford to lose.** Early software, no insurance, no recourse. That is not a
formality — it is the actual situation.

---

## 10. Glossary

**APR** — the yearly rate a borrower pays. Moves constantly.

**APY** — the yearly rate a lender earns. Moves constantly.

**Approve** — permission you grant a token contract allowing this platform to move a set amount of your
tokens. Required before a first deposit or repayment.

**Address** — your public identity on the blockchain, starting `0x`. Safe to share.

**Available liquidity** — USDC in the pool not currently lent out. The cap on withdrawals right now.

**Blockchain** — a shared, permanent record book kept by many computers.

**Collateral** — what you lock up to back a loan. Here, WETH.

**Decimals** — how many digits after the point a token uses. USDC 6, WETH 18.

**ERC-20** — the standard rulebook most tokens follow, which is why wallets and platforms handle them
uniformly.

**ETH** — the blockchain's own currency. Pays gas. Cannot be lent or used as collateral here.

**Gas** — the network's fee for processing a transaction, paid in ETH. Not ours.

**Health factor** — how safe your loan is. Above 1.00 fine, below 1.00 liquidatable, 1.50+ comfortable.

**Interest spread** — the 10% of borrower interest kept by the protocol to build a reserve, instead of going
to lenders. The only ongoing fee.

**Liquidation** — when your health factor falls below 1.00 and someone repays part of your debt in exchange
for your collateral plus a 5% bonus.

**Liquidation bonus** — the 5% reward paid to whoever resolves an unsafe position.

**Liquidation threshold** — 80%. The LTV at which liquidation becomes possible.

**LTV (loan-to-value)** — your debt as a percentage of your collateral's value. Lower is safer.

**Maximum borrow** — 75%. The most you may borrow against your collateral.

**Oracle / price feed** — the service publishing the WETH price on-chain, used to judge loan safety.

**Pool** — the shared pot of USDC that lenders fill and borrowers draw from.

**Shortfall** — when a liquidation pays out more than it recovers, because the collateral was worth less
than the debt plus bonus. Bad debt for the protocol.

**Stale price** — a price feed that has not updated recently enough to trust. Actions are disabled until it
recovers.

**Testnet** — a practice blockchain with free, worthless money. Where this platform currently runs.

**Token** — a unit of value on a blockchain that is not the blockchain's own currency.

**Transaction** — any action that changes the blockchain record. Needs your signature. Final once confirmed.

**USDC** — a token designed to hold a value of about one US dollar. The lending token here.

**Utilization** — the share of the pool currently lent out. Drives the rates. Rises steeply past 80%.

**Wallet** — the app holding the key that proves your money is yours. Not an account with us.

**WETH (Wrapped Ether)** — the ERC-20 token version of ETH, worth the same. The collateral token here.

---

## A closing note

This is a learning project built to real standards, not a finished commercial product. The contracts are
tested thoroughly, the rules are the same for everyone, and nothing about your money is hidden from you.

But it is early, it runs on a test network, and several screens are not connected yet. Treat it as a place
to genuinely understand how lending and borrowing on a blockchain works — which is worth doing, because the
knowledge transfers directly to platforms handling real money.

Start with **Welcome**. Then **Practice**. Then come back to this guide when something needs explaining.
