create table wallets
(
  id text not null primary key,
  user_id text not null,
  created_at timestamp with time zone,
  updated_at timestamp with time zone
);

alter table wallets owner to postgres;

create table balances
(
  id text not null primary key,
  wallet_id text not null constraint fk_wallets_balances references wallets,
  currency text not null,
  amount numeric(10, 2) not null,
  created_at timestamp with time zone,
  updated_at timestamp with time zone
);

alter table balances owner to postgres;

INSERT INTO public.wallets (id, user_id, created_at, updated_at) VALUES ('01HPMV01XPAXCG242W7SZWD0S5', '01HPMV01XPAXCG242W7SZWD0S7', '2024-08-20 16:48:36.965000 +00:00', '2024-08-20 16:48:38.153000 +00:00');
INSERT INTO public.balances (id, wallet_id, currency, amount, created_at, updated_at) VALUES ('01HPMV01XPAXCG242W7SZWD0S9', '01HPMV01XPAXCG242W7SZWD0S5', 'TRY', 30.10, '2024-08-20 16:52:26.735000 +00:00', '2024-08-20 16:52:27.807000 +00:00');
