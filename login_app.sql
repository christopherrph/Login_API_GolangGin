PGDMP         &                x         	   login_app    12.3    12.3                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16393 	   login_app    DATABASE     �   CREATE DATABASE login_app WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'English_United States.1252' LC_CTYPE = 'English_United States.1252';
    DROP DATABASE login_app;
                postgres    false            �            1259    16403    user    TABLE     �   CREATE TABLE public."user" (
    id_user integer NOT NULL,
    user_username character varying(100) NOT NULL,
    user_password character varying(250) NOT NULL,
    user_status character varying(50) NOT NULL
);
    DROP TABLE public."user";
       public         heap    postgres    false            �            1259    16401    user_id_user_seq    SEQUENCE     �   CREATE SEQUENCE public.user_id_user_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.user_id_user_seq;
       public          postgres    false    203                       0    0    user_id_user_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.user_id_user_seq OWNED BY public."user".id_user;
          public          postgres    false    202            
           2604    16406    user id_user    DEFAULT     n   ALTER TABLE ONLY public."user" ALTER COLUMN id_user SET DEFAULT nextval('public.user_id_user_seq'::regclass);
 =   ALTER TABLE public."user" ALTER COLUMN id_user DROP DEFAULT;
       public          postgres    false    202    203    203                      0    16403    user 
   TABLE DATA           T   COPY public."user" (id_user, user_username, user_password, user_status) FROM stdin;
    public          postgres    false    203          	           0    0    user_id_user_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.user_id_user_seq', 81, true);
          public          postgres    false    202            �
           2606    16408    user user_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id_user);
 :   ALTER TABLE ONLY public."user" DROP CONSTRAINT user_pkey;
       public            postgres    false    203               �   x�M�Ar�0 �u8k���Zš��`a����?�� zzW:�<2A�Zd�=�v�Y`Ӏ��Β�^����
�a)�֌c?o�{����]Է~��i���V�|\�%6��f�b8�;�v.OSF�Z�_}|?J�oXm�*�Ò��R�x�+��}�%D������7�v��{s�㍘E�'9$�*�iI�]j^ُcY��I'     