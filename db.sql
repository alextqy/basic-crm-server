PGDMP  4            	        |            BasicCrm    16.2    16.2 6    *           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            +           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            ,           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            -           1262    32768    BasicCrm    DATABASE     l   CREATE DATABASE "BasicCrm" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C';
    DROP DATABASE "BasicCrm";
                postgres    false            �            1259    32769    Admin    TABLE     �  CREATE TABLE public."Admin" (
    "ID" integer PRIMARY KEY NOT NULL,
    "Account" character varying(64) NOT NULL,
    "Password" character varying(64) NOT NULL,
    "Name" character varying(64) NOT NULL,
    "Level" smallint DEFAULT 1 NOT NULL,
    "Status" smallint DEFAULT 1 NOT NULL,
    "Remark" text,
    "Token" character varying(64),
    "CreationTime" integer DEFAULT 0 NOT NULL
);
    DROP TABLE public."Admin";
       public         heap    postgres    false            .           0    0 
   TABLE "Admin"    COMMENT     7   COMMENT ON TABLE public."Admin" IS '客服/管理员';
          public          postgres    false    215            /           0    0    COLUMN "Admin"."Account"    COMMENT     8   COMMENT ON COLUMN public."Admin"."Account" IS '账号';
          public          postgres    false    215            0           0    0    COLUMN "Admin"."Password"    COMMENT     9   COMMENT ON COLUMN public."Admin"."Password" IS '密码';
          public          postgres    false    215            1           0    0    COLUMN "Admin"."Name"    COMMENT     5   COMMENT ON COLUMN public."Admin"."Name" IS '名称';
          public          postgres    false    215            2           0    0    COLUMN "Admin"."Level"    COMMENT     W   COMMENT ON COLUMN public."Admin"."Level" IS '管理员等级 1普通 2超级 3临时';
          public          postgres    false    215            3           0    0    COLUMN "Admin"."Status"    COMMENT     M   COMMENT ON COLUMN public."Admin"."Status" IS '账号状态 1正常 2禁用';
          public          postgres    false    215            4           0    0    COLUMN "Admin"."Remark"    COMMENT     7   COMMENT ON COLUMN public."Admin"."Remark" IS '备注';
          public          postgres    false    215            �            1259    32841    Company    TABLE     �   CREATE TABLE public."Company" (
    "ID" integer PRIMARY KEY NOT NULL,
    "CompanyName" character varying(256) NOT NULL,
    "CreationTime" integer DEFAULT 0 NOT NULL,
    "Remark" text
);
    DROP TABLE public."Company";
       public         heap    postgres    false            5           0    0    TABLE "Company"    COMMENT     ;   COMMENT ON TABLE public."Company" IS '客户所属公司';
          public          postgres    false    221            6           0    0    COLUMN "Company"."Remark"    COMMENT     9   COMMENT ON COLUMN public."Company"."Remark" IS '备注';
          public          postgres    false    221            �            1259    32782    Customer    TABLE     �  CREATE TABLE public."Customer" (
    "ID" integer PRIMARY KEY NOT NULL,
    "Name" character varying(64) NOT NULL,
    "Birthday" integer DEFAULT 0 NOT NULL,
    "Gender" smallint DEFAULT 0 NOT NULL,
    "Email" character varying(256) NOT NULL,
    "Tel" character varying(64) NOT NULL,
    "CustomerInfo" text,
    "Priority" smallint DEFAULT 1 NOT NULL,
    "CreationTime" integer DEFAULT 0 NOT NULL,
    "CompanyID" integer DEFAULT 0 NOT NULL,
    "ManagerID" integer DEFAULT 0 NOT NULL
);
    DROP TABLE public."Customer";
       public         heap    postgres    false            7           0    0    TABLE "Customer"    COMMENT     6   COMMENT ON TABLE public."Customer" IS '顾客数据';
          public          postgres    false    216            8           0    0    COLUMN "Customer"."Gender"    COMMENT     L   COMMENT ON COLUMN public."Customer"."Gender" IS '性别 1女 2男 3其他';
          public          postgres    false    216            9           0    0     COLUMN "Customer"."CustomerInfo"    COMMENT     F   COMMENT ON COLUMN public."Customer"."CustomerInfo" IS '客户信息';
          public          postgres    false    216            :           0    0    COLUMN "Customer"."Priority"    COMMENT     P   COMMENT ON COLUMN public."Customer"."Priority" IS '优先级 1默认 99最高';
          public          postgres    false    216            ;           0    0     COLUMN "Customer"."CreationTime"    COMMENT     F   COMMENT ON COLUMN public."Customer"."CreationTime" IS '创建时间';
          public          postgres    false    216            <           0    0    COLUMN "Customer"."CompanyID"    COMMENT     C   COMMENT ON COLUMN public."Customer"."CompanyID" IS '关联企业';
          public          postgres    false    216            =           0    0    COLUMN "Customer"."ManagerID"    COMMENT     C   COMMENT ON COLUMN public."Customer"."ManagerID" IS '客服经理';
          public          postgres    false    216            �            1259    32808    Manager    TABLE     �  CREATE TABLE public."Manager" (
    "ID" integer PRIMARY KEY NOT NULL,
    "Account" character varying(64) NOT NULL,
    "Password" character varying(64) NOT NULL,
    "Name" character varying(64) NOT NULL,
    "Level" smallint DEFAULT 1 NOT NULL,
    "Status" smallint DEFAULT 1 NOT NULL,
    "Remark" text,
    "Token" character varying(64),
    "CreationTime" integer DEFAULT 0 NOT NULL,
    "GroupID" integer DEFAULT 0 NOT NULL
);
    DROP TABLE public."Manager";
       public         heap    postgres    false            >           0    0    TABLE "Manager"    COMMENT     5   COMMENT ON TABLE public."Manager" IS '销售经理';
          public          postgres    false    218            ?           0    0    COLUMN "Manager"."Level"    COMMENT     H   COMMENT ON COLUMN public."Manager"."Level" IS '等级 1普通 2主管';
          public          postgres    false    218            @           0    0    COLUMN "Manager"."Status"    COMMENT     O   COMMENT ON COLUMN public."Manager"."Status" IS '账号状态 1正常 2禁用';
          public          postgres    false    218            A           0    0    COLUMN "Manager"."Remark"    COMMENT     9   COMMENT ON COLUMN public."Manager"."Remark" IS '备注';
          public          postgres    false    218            �            1259    32821    ManagerGroup    TABLE     �   CREATE TABLE public."ManagerGroup" (
    "ID" integer PRIMARY KEY NOT NULL,
    "GroupName" character varying(128) NOT NULL,
    "CreationTime" integer DEFAULT 0 NOT NULL,
    "Remark" text
);
 "   DROP TABLE public."ManagerGroup";
       public         heap    postgres    false            B           0    0    TABLE "ManagerGroup"    COMMENT     7   COMMENT ON TABLE public."ManagerGroup" IS '销售组';
          public          postgres    false    219            C           0    0    COLUMN "ManagerGroup"."Remark"    COMMENT     >   COMMENT ON COLUMN public."ManagerGroup"."Remark" IS '备注';
          public          postgres    false    219            �            1259    32830 	   SalesPlan    TABLE     2  CREATE TABLE public."SalesPlan" (
    "ID" integer PRIMARY KEY NOT NULL,
    "PlanName" character varying(64) NOT NULL,
    "TargetID" integer DEFAULT 0 NOT NULL,
    "PlanContent" text,
    "CreationTime" integer DEFAULT 0 NOT NULL,
    "Status" smallint DEFAULT 1 NOT NULL,
    "Budget" money DEFAULT 0
);
    DROP TABLE public."SalesPlan";
       public         heap    postgres    false            D           0    0    TABLE "SalesPlan"    COMMENT     7   COMMENT ON TABLE public."SalesPlan" IS '销售计划';
          public          postgres    false    220            E           0    0    COLUMN "SalesPlan"."Status"    COMMENT     W   COMMENT ON COLUMN public."SalesPlan"."Status" IS '计划状态 1未完成 2已完成';
          public          postgres    false    220            �            1259    32797    SalesTarget    TABLE     B  CREATE TABLE public."SalesTarget" (
    "ID" integer PRIMARY KEY NOT NULL,
    "TargetName" character varying(64) NOT NULL,
    "ExpirationDate" integer DEFAULT 0 NOT NULL,
    "CreationTime" integer DEFAULT 0 NOT NULL,
    "AchievementRate" money DEFAULT 0 NOT NULL,
    "CustomerID" integer NOT NULL,
    "Remark" text
);
 !   DROP TABLE public."SalesTarget";
       public         heap    postgres    false            F           0    0    TABLE "SalesTarget"    COMMENT     9   COMMENT ON TABLE public."SalesTarget" IS '销售目标';
          public          postgres    false    217            G           0    0 %   COLUMN "SalesTarget"."ExpirationDate"    COMMENT     K   COMMENT ON COLUMN public."SalesTarget"."ExpirationDate" IS '截止日期';
          public          postgres    false    217            H           0    0 &   COLUMN "SalesTarget"."AchievementRate"    COMMENT     O   COMMENT ON COLUMN public."SalesTarget"."AchievementRate" IS '目标达成率';
          public          postgres    false    217            I           0    0 !   COLUMN "SalesTarget"."CustomerID"    COMMENT     G   COMMENT ON COLUMN public."SalesTarget"."CustomerID" IS '关联客户';
          public          postgres    false    217            J           0    0    COLUMN "SalesTarget"."Remark"    COMMENT     =   COMMENT ON COLUMN public."SalesTarget"."Remark" IS '备注';
          public          postgres    false    217            !          0    32769    Admin 
   TABLE DATA           |   COPY public."Admin" ("ID", "Account", "Password", "Name", "Level", "Status", "Remark", "Token", "CreationTime") FROM stdin;
    public          postgres    false    215   �8       '          0    32841    Company 
   TABLE DATA           R   COPY public."Company" ("ID", "CompanyName", "CreationTime", "Remark") FROM stdin;
    public          postgres    false    221   C9       "          0    32782    Customer 
   TABLE DATA           �   COPY public."Customer" ("ID", "Name", "Birthday", "Gender", "Email", "Tel", "CustomerInfo", "Priority", "CreationTime", "CompanyID", "ManagerID") FROM stdin;
    public          postgres    false    216   `9       $          0    32808    Manager 
   TABLE DATA           �   COPY public."Manager" ("ID", "Account", "Password", "Name", "Level", "Status", "Remark", "Token", "CreationTime", "GroupID") FROM stdin;
    public          postgres    false    218   }9       %          0    32821    ManagerGroup 
   TABLE DATA           U   COPY public."ManagerGroup" ("ID", "GroupName", "CreationTime", "Remark") FROM stdin;
    public          postgres    false    219   �9       &          0    32830 	   SalesPlan 
   TABLE DATA           v   COPY public."SalesPlan" ("ID", "PlanName", "TargetID", "PlanContent", "CreationTime", "Status", "Budget") FROM stdin;
    public          postgres    false    220   �9       #          0    32797    SalesTarget 
   TABLE DATA           �   COPY public."SalesTarget" ("ID", "TargetName", "ExpirationDate", "CreationTime", "AchievementRate", "CustomerID", "Remark") FROM stdin;
    public          postgres    false    217   �9       �           2606    32847    Company Company_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public."Company"
    ADD CONSTRAINT "Company_pkey" PRIMARY KEY ("ID");
 B   ALTER TABLE ONLY public."Company" DROP CONSTRAINT "Company_pkey";
       public            postgres    false    221            �           2606    32796    Customer Customer_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public."Customer"
    ADD CONSTRAINT "Customer_pkey" PRIMARY KEY ("ID");
 D   ALTER TABLE ONLY public."Customer" DROP CONSTRAINT "Customer_pkey";
       public            postgres    false    216            �           2606    32829    ManagerGroup ManagerGroup_pkey 
   CONSTRAINT     b   ALTER TABLE ONLY public."ManagerGroup"
    ADD CONSTRAINT "ManagerGroup_pkey" PRIMARY KEY ("ID");
 L   ALTER TABLE ONLY public."ManagerGroup" DROP CONSTRAINT "ManagerGroup_pkey";
       public            postgres    false    219            �           2606    32819    Manager Manager_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public."Manager"
    ADD CONSTRAINT "Manager_pkey" PRIMARY KEY ("ID");
 B   ALTER TABLE ONLY public."Manager" DROP CONSTRAINT "Manager_pkey";
       public            postgres    false    218            �           2606    32839    SalesPlan SalesPlan_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public."SalesPlan"
    ADD CONSTRAINT "SalesPlan_pkey" PRIMARY KEY ("ID");
 F   ALTER TABLE ONLY public."SalesPlan" DROP CONSTRAINT "SalesPlan_pkey";
       public            postgres    false    220            �           2606    32805    SalesTarget SalesTarget_pkey 
   CONSTRAINT     `   ALTER TABLE ONLY public."SalesTarget"
    ADD CONSTRAINT "SalesTarget_pkey" PRIMARY KEY ("ID");
 J   ALTER TABLE ONLY public."SalesTarget" DROP CONSTRAINT "SalesTarget_pkey";
       public            postgres    false    217            �           2606    32774    Admin admin_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public."Admin"
    ADD CONSTRAINT admin_pkey PRIMARY KEY ("ID");
 <   ALTER TABLE ONLY public."Admin" DROP CONSTRAINT admin_pkey;
       public            postgres    false    215            !   ?   x�3�LL����0O�LJ210H60361�4L3H5NJLL3L5JK1H�*3�4���4������ �/�      '   
   x������ � �      "   
   x������ � �      $   
   x������ � �      %   
   x������ � �      &   
   x������ � �      #   
   x������ � �     