PGDMP         +        
        x            test    12.4    12.4                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16393    test    DATABASE     �   CREATE DATABASE test WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'Chinese (Simplified)_China.936' LC_CTYPE = 'Chinese (Simplified)_China.936';
    DROP DATABASE test;
                postgres    false            �            1259    16443    todolist    TABLE     n   CREATE TABLE public.todolist (
    id integer NOT NULL,
    state text NOT NULL,
    journey text NOT NULL
);
    DROP TABLE public.todolist;
       public         heap    postgres    false            �            1259    16441 	   xc_id_seq    SEQUENCE     �   CREATE SEQUENCE public.xc_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
     DROP SEQUENCE public.xc_id_seq;
       public          postgres    false    203            	           0    0 	   xc_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.xc_id_seq OWNED BY public.todolist.id;
          public          postgres    false    202            �
           2604    16446    todolist id    DEFAULT     d   ALTER TABLE ONLY public.todolist ALTER COLUMN id SET DEFAULT nextval('public.xc_id_seq'::regclass);
 :   ALTER TABLE public.todolist ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    202    203    203                      0    16443    todolist 
   TABLE DATA           6   COPY public.todolist (id, state, journey) FROM stdin;
    public          postgres    false    203   A
       
           0    0 	   xc_id_seq    SEQUENCE SET     8   SELECT pg_catalog.setval('public.xc_id_seq', 37, true);
          public          postgres    false    202            �
           2606    16451    todolist xc_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.todolist
    ADD CONSTRAINT xc_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.todolist DROP CONSTRAINT xc_pkey;
       public            postgres    false    203               i   x�34�|�}��u=�:&p>���b�>.cS�gsVAŞ�kx�w��%ӞuN~>������s'��=���I�Ӿ�O�.{�`!��9�f�gs:��o��(����� HMA     